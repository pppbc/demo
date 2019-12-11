import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import cartGoods from './cartGoods.js'
Vue.use(Vuex)

export default new Vuex.Store({
	state:{
		totalBlock:0,	//总积分
		username:'',	//用户名
		token:localStorage.getItem("token"),	//token
		userID:'',	//用户ID
		avatar:'',	//用户头像
		nickname:'',	//用户昵称
		gender:'',		//用户性别
		email:'',		//用户邮箱
		birthday:'1900年1月1日',	//出生日期
		userinfo:'',		//用户信息
		footbarShow:true,	
		active:'home',	
		goods:'',
		allDatas:cartGoods,		//购物车列表
		receiverAddressList:[],	//收货地址列表
		currentReceiverAddress:'',	//当前收货地址
		blockArray:[],		//积分列表
		businessArray:[],	//商家列表
		allActicities:[],	//活动列表
		requestHomeGoods:1,	//home页商品请求页数
		requestActivities:1,	//活动请求页数
		requestMallcount:1,	//mall页商品请求页数
		requestAllCommentsCount:1,	//评价请求页数
		requestunpaidListCount:1,	//待付款列表请求页数
		requestUnshippedListCount:1,	//待发货列表请求页数
		requestunreceivedListCount:1,	//待收货列表请求页数
		requestCommentListCount:1,	//待评价列表请求页数
		requestAllOrdersListCount:1,	//全部订单请求页数
		requestRefundReturnListCount:1,	//退款退货列表请求页数
		allComments:[],		
		goodsDetail:[],
		fromPage:'',
		checkAll:true,
		confirmOrder:'',
		paymentData:'',
		orderActive:'待付款',
		orderDetail:'',
		commentData:'',
		refundReturnData:'',
		hasMessage:false,
		chatting:'',
	},
	mutations:{
		totalBlock(state,totalBlock){
			state.totalBlock = totalBlock
		},
		userinfo(state,data){
			state.userinfo = data
		},
		username(state,username){
			state.username = username
		},
		token(state,token){
			state.token = token
		},
		userID(state,userID){
			state.userID = userID
		},
		avatar(state,avatar){
			state.avatar = avatar
		},
		nickname(state,nickname){
			state.nickname = nickname
		},
		gender(state,gender){
			state.gender = gender
		},
		email(state,email){
			state.email = email
		},
		birthday(state,birthday){
			state.birthday = birthday
		},
		goodsDetail(state,data){
			state.goodsDetail = data
		},
		footbarShow(state){
			state.footbarShow = !state.footbarShow
		},
		changeTabbarActive(state,active){
			state.active = active
		},
		addGoodsData(state,data){
			state.allDatas.data.push(data)
		},
		updateCart(state,data){
			state.allDatas.data[data.length-1].Quantity = data.quantity
		},
		allDatasData(state,data){
			state.allDatas.data = data
		},
		addToCart1(state,data){
			state.allDatas.data.push({
				UserID:state.userID,
				ID:data.data.ID,
				MainImage:data.data.MainImage,
				ProductName:data.data.Name,
				Checked:true,
				Price:data.data.Price,
				Quantity:data.num,
			})
		},
		addToCart2(state,data){
			state.allDatas.data[data.index].Quantity += data.num
		},
		deleteFromCartLocal(state,index){
			state.allDatas.data.splice(index,1)
		},
		checkAll(state,value){
			state.checkAll = value
		},
		deleteGoodFromCart(state){
			state.allDatas.data = state.allDatas.data.filter(val=>{
				return	val.Checked === false
			})
		},
		receiverAddressList(state,data){
			state.receiverAddressList = data
		},
		currentReceiverAddress(state,index){
			if (index >= 0) {
				state.currentReceiverAddress = state.receiverAddressList[index]
			} else{
				for (let item of state.receiverAddressList) {
					if (item.DefaultAddress === 1) {
						state.currentReceiverAddress = item
					}
				}
			}
		},
		saveResult(state,data){
			state.blockArray = data.filter(val=>{
				return val.Balance !== 0
			})
		},
		businessArray(state,data){
			state.businessArray = data
		},
		allActicities(state,data){
			state.allActicities = [...state.allActicities,...data]
		},
		requestHomeGoods(state,count){
			if (count === 0) {
				state.requestHomeGoods = 1
			} else{
				state.requestHomeGoods += 1
			}
		},
		requestActivities(state,count){
			if (count === 0) {
				state.requestActivities = 1
			} else{
				state.requestActivities += 1
			}
		},
		requestMallcount(state,count){
			if (count === 0) {
				state.requestMallcount = 1
			} else{
				state.requestMallcount += 1
			}
		},
		requestAllCommentsCount(state,count){
			if (count === 0) {
				state.requestAllCommentsCount = 1
			} else{
				state.requestAllCommentsCount += 1
			}
		},
		requestunpaidListCount(state,count){
			if (count === 0) {
				state.requestunpaidListCount = 1
			} else{
				state.requestunpaidListCount += 1
			}
		},
		requestUnshippedListCount(state,count){
			if (count === 0) {
				state.requestUnshippedListCount = 1
			} else{
				state.requestUnshippedListCount += 1
			}
		},
		requestunreceivedListCount(state,count){
			if (count === 0) {
				state.requestunreceivedListCount = 1
			} else{
				state.requestunreceivedListCount += 1
			}
		},
		requestCommentListCount(state,count){
			if (count === 0) {
				state.requestCommentListCount = 1
			} else{
				state.requestCommentListCount += 1
			}
		},
		requestAllOrdersListCount(state,count){
			if (count === 0) {
				state.requestAllOrdersListCount = 1
			} else{
				state.requestAllOrdersListCount += 1
			}
		},
		requestRefundReturnListCount(state,count){
			if (count === 0) {
				state.requestRefundReturnListCount = 1
			} else{
				state.requestRefundReturnListCount += 1
			}
		},
		allComments(state,data){
			if (data) {
				state.allComments = [...state.allComments,...data]
			} else{
				state.allComments = []
			}
		},
		exchangeReceiverItem(state,index){
			for (let item of state.receiverAddressList) {
				item.DefaultAddress = -1
			}
			state.receiverAddressList[index].DefaultAddress = 1
		},
		deleteReceiverItem(state,index){
			state.receiverAddressList.splice(index,1)
		},
		addReceiverItem(state,data){
			state.receiverAddressList.push(data)
		},
		modifyReceiverAddressList (state,item) { //修改收货地址
			state.receiverAddressList[parseInt(item.index)] = item.data
		},
		fromPage(state,page){
			state.fromPage = page
		},
		confirmOrder(state,data){
			state.confirmOrder = data
		},
		paymentData(state,data){
			state.paymentData = data
		},
		orderActive(state,data){
			state.orderActive = data
		},
		orderDetail(state,data){
			state.orderDetail = data
		},
		commentData(state,data){
			state.commentData = data
		},
		refundReturnData(state,data){
			state.refundReturnData = data
		},
		hasMessage(state,data){
			state.hasMessage = data
		},
		chatting(state,data){
			state.chatting = data
		}
	},
	getters:{
		allBlock(state){
			var totalBlock = 0
			for (var i = 0; i < state.blockArray.length; i++) {
				var e = state.blockArray[i]
				totalBlock += e.Balance * e.PointFate
			}
			totalBlock += state.totalBlock
			return totalBlock
		}
	},
	plugins:[createPersistedState()]
})