<template>
	<div>
		<van-nav-bar
			title="退款/退货"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		
		<van-tabs v-model="active" color="#ff0000" title-active-color="#ff0000" @change="tabChange" animated sticky :offset-top="offsetTop">
			<van-tab v-for="(item,index) in tabs" :key="index" :title="item" :name="item">
				<van-list
					v-model="loading"
					:finished="finished"
					finished-text="没有更多了"
					@load="onLoad"
				>
					<order-cell
						v-for="(item,index) in datalist" 
						:key="index"
						:title="orderStatus(item.Status)"
						:time="item.CreateTime.replace('T',' ').replace('Z',' ')"
						:image="item.ProductImage"
						:name="item.ProductName"
						:goodsNum="item.Quantity"
						:totalPrice="item.TotalPrice"
						:orderNo="item.OrderNo"
						@orderDetail="orderDetail(item)"
						@cancelRefund="cancelRefund(item,index)"
					></order-cell>
				</van-list>
			</van-tab>
		</van-tabs>
		<router-view></router-view>
	</div>
</template>

<script>
	import {NavBar,Tabs,Tab,List,Toast,Dialog} from 'vant'
	import {Images} from '@/resources/index.js'
	import orderCell from '@/components/orderCell.vue'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Tabs.name]:Tabs,
			[Tab.name]:Tab,
			[List.name]:List,
			[Toast.name]:Toast,
			[Dialog.Component.name]:Dialog.Component,
			orderCell
		},
		data(){
			return {
				bg:Images.indent.ic_bg,
				active:'退款',
				datalist:'',
				tabs:['退款','退货'],
				title:'',
				loading:false,
				finished:false,
				offsetTop:46
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive','mine')
				this.$router.push('/main/mine')
			},
			tabChange(name){
				this.loading = false,
				this.finished = false
				this.datalist = []
				this.$store.commit('requestRefundReturnListCount',0)
				this.getList(this.$store.state.requestRefundReturnListCount)
			},
			getList(count){
				var url = ''
				if (this.active === '退款') {
					url = '/order/orderReturnList?user_id='+this.$store.state.userID+'&token='+this.$store.state.token
				} else if (this.active === '退货') {
					url = '/order/orderProductList?user_id='+this.$store.state.userID+'&token='+this.$store.state.token
				} else{
					return
				}
				this.axios.get(url).then((response)=>{
					console.log(response)
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						if (response.data.obj.length > 0) {
							this.$store.commit('requestRefundReturnListCount',1)
							this.datalist = [...this.datalist,...response.data.obj]
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
					this.getList(this.$store.state.requestRefundReturnListCount)
					this.loading = false
					this.finished = true
				},500)
			},
			orderDetail(item){
				item.title = this.active
				this.$store.commit('orderDetail',item)
				this.$router.push('/orderDetail')
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
								this.datalist.splice(index,1)
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
				return function(status){
					var title = ''
					if (this.active === '退款') {
						if (status === 1) {
							title = '退款中'
						}
						if (status === 2) {
							title = '退款成功'
						}
						if (status === 3) {
							title = '退款失败'
						}
					}
					if (this.active === '退货') {
						if (status === 1) {
							title = '退货申请中'
						}
						if (status === 2) {
							title = '退货申请成功'
						}
						if (status === 3) {
							title = '退货申请失败'
						}
						if (status === 4) {
							title = '等待卖家收货'
						}
						if (status === 5) {
							title = '卖家同意退款'
						}
						if (status === 6) {
							title = '卖家拒绝退款'
						}
					}
					return title
				}
			}
		},
		beforeMount() {
			this.active = this.$route.params.name
		}
	}
</script>

<style scoped>
	.van-nav-bar{
		background: url(../../assets/Images/Mine/background1.png);
	}
	.van-nav-bar__title{
		color: white;
	}
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: white;
	}
	.van-tabs{
		margin-top: 2.875rem;
	}
</style>
