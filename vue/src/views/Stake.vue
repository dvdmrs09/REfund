<template>
  <div class="child-container">
    <div v-if="store.stakePopup">
      <StakePopup :manage="manageStake"/>
    </div>
    <div v-if="typeof(rows_rewards) != 'undefined' && rows_rewards.length > 0" id="your-stake-div">
      <header class="funds-header">
        <h2 class="title">Your Stake</h2>
      </header>
      <div class="grid-div">
        <v-grid
          id="rewards-grid"
          style="min-height: 300px;"
          theme="material"
          :source="rows_rewards"
          :columns="columns_rewards"
          filter=true
          readonly=true
          resize=true
          col-size=200
          row-size=60
          can-focus=false
          row-headers=true
          @beforecellfocus="e => onRowClick(e, true)"
        ></v-grid>
      </div>
    </div>
    <div style="margin-top: 25px;" id="validators-div">
      <header class="funds-header">
        <h2 class="title">All Validators</h2>
      </header>
      <div class="grid-div">
        <v-grid
          theme="material"
          :source="store.validators"
          :columns="columns"
          filter=true
          readonly=true
          resize=true
          col-size=200
          row-size=60
          can-focus=false
          row-headers=true
          @beforecellfocus="e => onRowClick(e, false)"
        ></v-grid>
      </div>
    </div>
  </div>
</template>
 
<script>
import VGrid, {VGridVueTemplate} from "@revolist/vue3-datagrid";
import StakeButton from '../components/StakeButton.vue'
import StakePopup from '../components/StakePopup.vue'
import { computed } from 'vue';
import { useStore } from 'vuex';
import { SpTheme } from '@starport/vue';
import flatten from 'flat';
import _ from 'lodash';
import { store } from '../store/local/store.js';

export default {
  name: "Stake",
  beforeMount() {

    let $s = useStore()

    let address = computed(() => {
        return $s.getters['common/wallet/address']
    })

    if(address) {

      $s.dispatch("cosmos.staking.v1beta1/QueryDelegatorValidators", {params: { delegator_addr: address.value }, subscribe: true, all: false })

      $s.dispatch("cosmos.staking.v1beta1/QueryDelegatorDelegations", {params: { delegator_addr: address.value }, subscribe: true, all: false })

      $s.dispatch("cosmos.distribution.v1beta1/QueryDelegationTotalRewards", {params: { delegator_address: address.value }, subscribe: true, all: false })
    
    }

  },
  beforeUpdate() {

    if(this.updateData) {

      let $s = useStore()

      let address = computed(() => {
          return $s.getters['common/wallet/address']
      })

      if(address) {

        $s.dispatch("cosmos.staking.v1beta1/QueryDelegatorValidators", {params: { delegator_addr: address.value }, subscribe: true, all: false })

        $s.dispatch("cosmos.staking.v1beta1/QueryDelegatorDelegations", {params: { delegator_addr: address.value }, subscribe: true, all: false })

        $s.dispatch("cosmos.distribution.v1beta1/QueryDelegationTotalRewards", {params: { delegator_address: address.value }, subscribe: true, all: false })
      
        this.updateData = false
      }

    }

  },
  data() {
    let $s = useStore()

    let address = computed(() => {
        return $s.getters['common/wallet/address']
    })

    let vals = computed(() => {
      var validators_raw = $s["getters"]["cosmos.staking.v1beta1/getValidators"]()
      var validators = validators_raw["validators"]
      validators = _.orderBy(validators, [function(o) { return Number(o.tokens); }], ["desc"]);
      if (typeof(validators) != "undefined") {
        for (let i = 0; i < validators.length; i++) {
          validators[i]["tokens"] = String(_.round(Number(validators[i]["tokens"])/1000000, 2)) + " FETF"
          validators[i] = flatten(validators[i]);
          validators[i]["commission.commission_rates.rate"] = String(_.round(Number(validators[i]["commission.commission_rates.rate"]) * 100, 2)) + "%"
        }
      }
      if(typeof(validators_raw) == "undefined") { this.total = Number(validators_raw["pagination"]["total"]) }
      return validators
    })

    let rewards = computed(() => {
      var validators_raw = $s["getters"]["cosmos.staking.v1beta1/getDelegatorValidators"]({params: { delegator_addr: address.value }})
      var validators = validators_raw["validators"]
      if (typeof(validators) != "undefined") {
        for (let i = 0; i < validators.length; i++) {
          validators[i] = flatten(validators[i]);
          // Get the all the rewards for curret wallet (from store)
          var all_rewards = $s["getters"]["cosmos.distribution.v1beta1/getDelegationTotalRewards"]({params: { delegator_address: address.value }})
          // Filter all the rewards for the current validator in the for loop
          var current_rewards = _.filter(all_rewards["rewards"], function(o) { return o.validator_address == validators[i]["operator_address"] })[0]
          // If no current rewards are found for validator, skip the current and continue the loop
          if(typeof(current_rewards) == "undefined") { continue }
          // Filter the rewards for just ufetf for teh validator found above
          var rewards = _.filter(current_rewards["reward"], function(o) { return o.denom == "ufetf" })[0]
          // If no base denom (fetf) rewards are found for validator, skip the current and continue the loop
          if(typeof(rewards) == "undefined") { continue }
          rewards["rewards.amount"] = String(_.round(Number(rewards["amount"])/1000000, 2)) + " FETF"
          rewards["rewards.denom"] = "fetf"
          if (rewards["amount"] == "NaN FETF") { rewards["amount"] = "0" + " FETF" }
          // Get the all the delegations for curret wallet (from store)
          var all_delegations = $s["getters"]["cosmos.staking.v1beta1/getDelegatorDelegations"]({params: { delegator_addr: address.value }})
          // Filter all the delegations for the current validator in the for loop
          var current_delegation = _.filter(all_delegations["delegation_responses"], function(o) { return o.delegation.validator_address == validators[i]["operator_address"] })[0]
          // If no current rewards are found for validator, skip the current and continue the loop
          if(typeof(current_delegation) == "undefined") { continue }
          current_delegation["data"] = {}
          current_delegation["data"]["delegation.amount"] = String(_.round(Number(current_delegation["balance"]["amount"])/1000000, 2)) + " FETF"
          current_delegation["data"]["delegation.denom"] = "fetf"
          validators[i] = {...validators[i], ...rewards, ...current_delegation.data}
        }
      }
      return typeof(validators) != "undefined" ? validators : []
    })

    return {
      columns: [
        { name: "Moniker", prop: "description.moniker"},
        { name: "Website", prop: "description.website"},
        { name: "Voting Power", prop: "tokens_string"},
        { name: "Commission", prop: "commission.commission_rates.rate_string" }
      ],
      columns_rewards: [
        { name: "Moniker", prop: "description.moniker"},
        { name: "Website", prop: "description.website"},
        { name: "Delegations", prop: "delegation.amount"},
        { name: "Rewards", prop: "rewards.amount"}
      ],
      rows_rewards: rewards,
      rows: vals,
      store: store,
      s: $s,
      page: 0,
      total: null,
      address,
      updateData: true
    };
  },
  methods: {
    beforeFocus(e) {
      e.preventDefault();
    },
    onRowClick(e, manageStake = false) {
      this.openManagePopup(manageStake, e.detail.model)
    },
    openManagePopup(manageStake, currentValidator) {
      if (store.stakePopup == false) {
        store.stakePopup = true
        store.currentValidator = currentValidator
        store.manageStake = manageStake
      } else {
        store.stakePopup = false
      }
    }
  },
  watch: {
    address(newAddr, oldAddr) {
      this.updateData = true
      if(newAddr) {

        this.s.dispatch("cosmos.staking.v1beta1/QueryDelegatorValidators", {params: { delegator_addr: newAddr }, subscribe: true, all: false })

        this.s.dispatch("cosmos.staking.v1beta1/QueryDelegatorDelegations", {params: { delegator_addr: newAddr }, subscribe: true, all: false })

        this.s.dispatch("cosmos.distribution.v1beta1/QueryDelegationTotalRewards", {params: { delegator_address: newAddr }, subscribe: true, all: false })
      
        this.s.dispatch("cosmos.staking.v1beta1/QueryValidators", {subscribe: true, all: false})

        this.updateData = false
      }
    }
  },
  components: {
    VGrid, SpTheme, StakeButton, StakePopup
  }
};
</script>

<style>
  revo-grid {
    min-height: 500px;
  }
  .title {
    font-family: Inter, serif;
    font-style: normal;
    font-weight: 600;
    font-size: 28px;
    line-height: 127%;
    /* identical to box height, or 36px */

    letter-spacing: -0.02em;
    font-feature-settings: 'zero';

    color: #000000;
    margin-top: 0;
  }
  revo-grid[theme=material] {
      font-family: "Inter", sans-serif;
  }
  revo-grid[theme=material] revogr-data .rgRow {
    line-height: 60px;
  }
  revo-grid[theme=material] revogr-data .rgRow:hover {
      background-color: rgba(233, 234, 237, 0.5);
      cursor: pointer
  }
  .child-container {
    padding-right: 32px;
    padding-left: 32px;
    padding-bottom: 32px;
  }
</style>

