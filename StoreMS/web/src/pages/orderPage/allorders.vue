<template>
	<div class="allorders">
		<van-list
			v-model="loading"
			:finished="finished"
			finished-text="没有更多了"
			@load="onLoad"
		>
			<order-cell 
				v-for="(item,index) in data" 
				:key="index"
				:title="orderStatus(item.Status)"
				:time="item.CreateTime.replace('T',' ').replace('Z',' ')"
				:image="item.ProductImage"
				:name="item.ProductName"
				:goodsNum="item.Quantity"
				:totalPrice="item.TotalPrice"
				:orderNo="item.OrderNo"
				@orderDetail="orderDetail(item)"
				@cancelOrder="cancelOrder(item,index)"
				@refund="refund(item)"
				@confirmReceipt="confirmReceipt(item,index)"
				@returnOrder="returnOrder(item)"
				@delete="deleteOrder(item,index)"
				@cancelRefund="cancelRefund(item,index)"
			></order-cell>
		</van-list>
	</div>
</template>

<script>
	import {Toast,List,Dialog} from 'vant'
	import orderCell from '@/components/orderCell.vue'
	export default{
		components:{
			[Toast.name]:Toast,
			[List.name]:List,
			[Dialog.Component.name]:Dialog.Component.name,
			orderCell
		},
		data(){
			return {
				data:[],
				title:'',
				loading:false,
				finished:false,
			}
		},
		methods:{
			getAllOrdersList(count){
				var url = '/order/allList?user_id='+this.$store.state.userID+'&page='+count+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					console.log(response)
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						if (response.data.obj.length>0) {
							this.$store.commit('requestAllOrdersListCount',1)
							this.data = [...this.data,...response.data.obj]
						} else{
							this.finished = true
						}
					} else{
						Toast('数据请求失败，请重试')
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			onLoad(){
				setTimeout(()=>{
					this.getAllOrdersList(this.$store.state.requestAllOrdersListCount)
					this.loading = false
				},1000)
			},
			orderDetail(item){
				this.$store.commit('orderDetail',item)
				this.$router.push('/orderDetail')
			},
			cancelOrder(item,index){
				Dialog.confirm({
					title:'提示',
					message:'是否确认取消订单'
				}).then(()=>{
					var url = '/order/cancelOrder?order_no='+item.OrderNo+'&user_id='+item.UserID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast('订单已取消')
								this.data.splice(index,1)
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			},
			refund(item){
				this.$store.commit('refundReturnData',item)
				this.$router.push('/refundApply/退款申请')
			},
			confirmReceipt(item,index){
				Dialog.confirm({
					title:'提示',
					message:'是否确认收货'
				}).then(()=>{
					var url = '/order/ensureOrder?order_no='+item.OrderNo+'&user_id='+item.UserID+'&product_id='+item.ProductID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							this.data.splice(index,1)
							Toast('已确认收货')
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			},
			returnOrder(item){
				this.$store.commit('refundReturnData',item)
				this.$router.push('/returnApply/退货申请')
			},
			deleteOrder(item,index){
				Dialog.confirm({
					title:'提示',
					message:'是否确认删除'
				}).then(()=>{
					var url = '/order/cancelOrder?order_no='+item.OrderNo+'&user_id='+item.UserID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast('订单已取消')
								this.data.splice(index,1)
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			},
			cancelRefund(item,index){
				Dialog.confirm({
					title:'提示',
					message:'是否取消退款'
				}).then(()=>{
					var url = '/order/cancelReturn?user_id='+this.$store.state.userID+'&order_no='+item.OrderNo+'&product_id='+item.ProductID+'&return_id='+item.ID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast(response.data.detail)
								this.data.splice(index,1)
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			}
		},
		computed:{
			orderStatus(){
				return function (status) {
					var title = ''
					if (status === 0) {
						title = '交易取消'
					}
					if (status === 10) {
						title = '待付款'
					}
					if (status === 20) {
						title = '待发货'
					}
					if (status === 30) {
						title = '待收货'
					}
					if (status === 40) {
						title = '交易完成'
					}
					if (status === 50) {
						title = '交易关闭'
					}
					if (status === 21) {
						title = '退款中'
					}
					if (status === 22) {
						title = '退款成功'
					}
					if (status === 23) {
						title = '退款失败'
					}
					return title
				}
			}
		},
		beforeRouteEnter(to,from,next) {
			next(vm=>{
				vm.$store.commit('requestAllOrdersListCount',0)
				vm.getAllOrdersList(vm.$store.state.requestAllOrdersListCount)
			})
		}
	}
</script>

<style scoped>
	.allorders{
		margin-top: 7.5rem;
		
	}
</style>
