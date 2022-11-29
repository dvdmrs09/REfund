import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { BaseDenom } from "defund-labs-defund-client-ts/defundlabs.defund.etf/types"
import { FundPrice } from "defund-labs-defund-client-ts/defundlabs.defund.etf/types"
import { Balances } from "defund-labs-defund-client-ts/defundlabs.defund.etf/types"
import { Holding } from "defund-labs-defund-client-ts/defundlabs.defund.etf/types"
import { Fund } from "defund-labs-defund-client-ts/defundlabs.defund.etf/types"
import { AddressMap } from "defund-labs-defund-client-ts/defundlabs.defund.etf/types"


export { BaseDenom, FundPrice, Balances, Holding, Fund, AddressMap };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Fund: {},
				FundAll: {},
				FundPrice: {},
				
				_Structure: {
						BaseDenom: getStructure(BaseDenom.fromPartial({})),
						FundPrice: getStructure(FundPrice.fromPartial({})),
						Balances: getStructure(Balances.fromPartial({})),
						Holding: getStructure(Holding.fromPartial({})),
						Fund: getStructure(Fund.fromPartial({})),
						AddressMap: getStructure(AddressMap.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getFund: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Fund[JSON.stringify(params)] ?? {}
		},
				getFundAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FundAll[JSON.stringify(params)] ?? {}
		},
				getFundPrice: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FundPrice[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: defundlabs.defund.etf initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryFund({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundEtf.query.queryFund( key.symbol)).data
				
					
				commit('QUERY', { query: 'Fund', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFund', payload: { options: { all }, params: {...key},query }})
				return getters['getFund']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFund API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFundAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundEtf.query.queryFundAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DefundlabsDefundEtf.query.queryFundAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'FundAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFundAll', payload: { options: { all }, params: {...key},query }})
				return getters['getFundAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFundAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFundPrice({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundEtf.query.queryFundPrice(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DefundlabsDefundEtf.query.queryFundPrice({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'FundPrice', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFundPrice', payload: { options: { all }, params: {...key},query }})
				return getters['getFundPrice']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFundPrice API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgRedeem({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundEtf.tx.sendMsgRedeem({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRedeem:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRedeem:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateFund({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundEtf.tx.sendMsgCreateFund({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateFund:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateFund:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundEtf.tx.sendMsgCreate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgRedeem({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundEtf.tx.msgRedeem({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRedeem:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRedeem:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateFund({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundEtf.tx.msgCreateFund({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateFund:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateFund:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundEtf.tx.msgCreate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreate:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
