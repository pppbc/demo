<template>
	<div>
		<van-nav-bar
			title="购物车"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		
		<van-row class="main">
			<van-col class="content" span="24" v-for="(item,index) in cartList" :key="index">
				<van-row>
					<van-col class="check" span="2">
						<van-checkbox v-model="item.Checked" @change="change(item,index)" checked-color="#ff0000"></van-checkbox>
					</van-col>
					<van-col class="img" span="8">
						<van-image :src="item.MainImage" width="5rem" height="5rem"/>
					</van-col>
					<van-col class="right" span="14">
						<van-row class="prdname">
							<van-col span="24">{{item.ProductName}}</van-col>
						</van-row>
						<van-row class="price">
							<van-col span="14">Block:{{item.Price*item.Quantity}}</van-col>
							<van-col span="10">x{{item.Quantity}}</van-col>
						</van-row>
						<van-row class="opera">
							<van-col span="14">
								<van-stepper v-model="item.Quantity" disable-input @change="change(item,index)"></van-stepper>
							</van-col>
							<van-col span="10">
								<van-button type="danger" @click="deleteFromCart(item.ID,index)">删除</van-button>
							</van-col>
						</van-row>
					</van-col>
				</van-row>
			</van-col>
		</van-row>
		
		<van-submit-bar
			:price="totalPrice"
			currency="B"
			button-text="去结算"
			@submit="onSubmit"
			safe-area-inset-bottom
		>
			<van-checkbox v-model="checkAll" @click="checkedAll" checked-color="#ff0000">全选</van-checkbox>
		</van-submit-bar>
		
	</div>
</template>

<script>
	import {NavBar,Checkbox,Image,Stepper,Button,Row,Col,SubmitBar,Toast} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Checkbox.name]:Checkbox,
			[Image.name]:Image,
			[Stepper.name]:Stepper,
			[Button.name]:Button,
			[Row.name]:Row,
			[Col.name]:Col,
			[SubmitBar.name]:SubmitBar,
			[Toast.name]:Toast
		},
		data(){
			return {
				checkAll:true,
				cartList:[],
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive',this.$store.state.fromPage)
				this.$router.back()
			},
			onSubmit(){
				var hasChecked = this.$store.state.allDatas.data.every(val=>{
					return val.Checked === false
				})
				if (!hasChecked) {
					this.$router.push('/confirmOrder')
				}else{
					Toast('还没有选择商品')
				}
			},
			change(item,index){
				var data ={
					UserID:this.$store.state.userID,
					ProductID:item.ProductID,
					Quantity:item.Quantity,
					Checked:item.Checked
				}
				var url = '/car/update?token='+this.$store.state.token
				this.axios.post(url,data).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						for (let val of this.cartList) {
							if (val['Checked']) {
								this.checkAll = true
								this.$store.commit('checkAll',this.checkAll)
							}else{
								this.checkAll = false
								this.$store.commit('checkAll',this.checkAll)
							}
						}
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			getUserCartDetail(){
				var url = '/cart/get?user_id='+this.$store.state.userID+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else	if(response.data.status === 1) {
						this.cartList = response.data.info
						this.$store.commit('allDatasData',response.data.info)
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			checkedAll(){
				this.checkAll = !this.checkAll
				if (this.checkAll) {
					for (let val of this.cartList) {
						val['Checked'] = true
					}
				} else{
					for (let val of this.cartList) {
						val['Checked'] = false
					}
				}
				var url = '/product/checkAll?user_id='+this.$store.state.userID+'&state='+this.checkAll+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					this.$store.commit('checkAll',this.checkAll)
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			deleteFromCart(id,index){
				var url = '/product/deleteProduct?id='+id+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					} else if(response.data.status === 1){
						this.$store.commit('deleteFromCartLocal',index)
						Toast('删除成功')
					} else {
						Toast('删除失败，请重试')
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			}
		},
		beforeMount(){
			this.$store.commit('allDatasData',[])
			this.checkAll = this.$store.state.checkAll
			this.getUserCartDetail()
		},
		computed:{
			totalPrice(){
				var totalPrice = 0
				for (let item of this.cartList) {
					if (item.Checked) {
						totalPrice += (item.Price * item.Quantity)*100
					}
				}
				return totalPrice
			}
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: #000000;
	}
	.main{
		margin: 3.125rem 0.625rem;
	}
	.content{
		background: white;
		border-radius: 0.5rem;
		overflow: hidden;
		margin-bottom: 0.625rem;
		height: 6.25rem;
	}
	
	.check .van-checkbox{
		margin: 2.5rem 0 0 0.5rem;
	}
	.img .van-image{
		margin-top: 0.625rem;
	}
	.prdname{
		text-align: left;
		font-size: 0.75rem;
		color: #999999;
		padding: 0.625rem 0.625rem 0 0;
	}
	.prdname .van-col{
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}
	.price{
		text-align: left;
		font-size: 0.875rem;
		padding: 0.625rem 0.625rem 0 0;
	}
	.price .van-col:first-child{
		color: red;
	}
	.price .van-col:last-child{
		text-align: center;
		color: #999999;
	}
	.opera{
		padding: 0.625rem 0.625rem 0 0;
	}
	.opera .van-button{
		height: 1.75rem;
		line-height: 1.75rem;
	}
	.van-submit-bar .van-checkbox{
		margin-left: 0.625rem;
	}
</style>
