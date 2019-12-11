<template>
	<div class="unreceived">
		<van-list
			v-model="loading"
			:finished="finished"
			finished-text="没有更多了"
			@load="onLoad"
		>
			<order-cell 
				v-for="(item,index) in data" 
				:key="index"
				title="待收货"
				:time="item.CreateTime.replace('T',' ').replace('Z',' ')"
				:image="item.ProductImage"
				:name="item.ProductName"
				:goodsNum="item.Quantity"
				:totalPrice="item.TotalPrice"
				:orderNo="item.OrderNo"
				@orderDetail="orderDetail(item)"
				@returnOrder="returnOrder(item)"
				@confirmReceipt="confirmReceipt(item,index)"
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
				loading:false,
				finished:false,
			}
		},
		methods:{
			getUnreceivedList(count){
				var url = '/order/sendList?user_id='+this.$store.state.userID+'&page='+count+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						if (response.data.obj.length>0) {
							this.$store.commit('requestunreceivedListCount',1)
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
					this.getUnreceivedList(this.$store.state.requestunreceivedListCount)
					this.loading = false
				},1000)
			},
			returnOrder(item){
				this.$store.commit('refundReturnData',item)
				this.$router.push('/returnApply/退货申请')
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
			orderDetail(item){
				this.$store.commit('orderDetail',item)
				this.$router.push('/orderDetail')
			}
		},
		beforeRouteEnter(to,from,next) {
			next(vm=>{
				vm.$store.commit('requestunreceivedListCount',0)
				vm.getUnreceivedList(vm.$store.state.requestunreceivedListCount)
			})
		}
	}
</script>

<style scoped>
	.unreceived{
		margin-top: 7.5rem;
		
	}
</style>
