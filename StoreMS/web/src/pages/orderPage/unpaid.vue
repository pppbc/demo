<template>
	<div class="unpaid">
		<van-list
			v-model="loading"
			:finished="finished"
			finished-text="没有更多了"
			@load="onLoad"
		>
			<order-cell 
				v-for="(item,index) in data" 
				:key="index"
				title="待付款"
				:time="item.CreateTime.replace('T',' ').replace('Z',' ')"
				:image="item.ProductImage"
				:name="item.ProductName"
				:goodsNum="item.Quantity"
				:totalPrice="item.TotalPrice"
				:orderNo="item.OrderNo"
				@cancelOrder="cancelOrder(item,index)"
				@orderDetail="orderDetail(item)"
			></order-cell>
		</van-list>
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show"></van-loading>
	</div>
</template>

<script>
	import {Toast,List,Dialog,Overlay,Loading} from 'vant'
	import orderCell from '@/components/orderCell.vue'
	export default{
		components:{
			[Toast.name]:Toast,
			[List.name]:List,
			[Dialog.Component.name]:Dialog.Component,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading,
			orderCell
		},
		data(){
			return {
				data:[],
				loading:false,
				finished:false,
				show:false,
			}
		},
		methods:{
			getUnpaidList(count){
				var url = '/order/unpayList?user_id='+this.$store.state.userID+'&page='+count+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						if (response.data.obj.length>0) {
							this.$store.commit('requestunpaidListCount',1)
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
					this.getUnpaidList(this.$store.state.requestunpaidListCount)
					this.loading = false
				},1000)
			},
			cancelOrder(item,index){
				Dialog.confirm({
					title:'提示',
					message:'是否确认取消订单'
				}).then(()=>{
					this.show = true
					var url = '/order/cancelOrder?order_no='+item.OrderNo+'&user_id='+item.UserID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast('订单已取消')
								this.show = false
								this.data.splice(index,1)
							}, 1000);
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
				vm.$store.commit('requestunpaidListCount',0)
				vm.getUnpaidList(vm.$store.state.requestunpaidListCount)
			})
		}
	}
</script>

<style scoped>
	.unpaid{
		margin-top: 7.5rem;
	}
	.overlay{
		background: rgba(0,0,0,0);
	}
	.van-loading{
		position: absolute;
		left: 45%;
		top: 45%;
	}
</style>
