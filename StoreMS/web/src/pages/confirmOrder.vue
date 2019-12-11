<template>
	<div>
		<van-nav-bar
			title="确认订单"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		<div class="content">
			<van-contact-card v-if="!currentAddr" @click="selectAddr"/>
			<div v-else>
				<van-row class="address" @click="selectAddr">
					<van-col span="4" class="child1">
						<van-icon :name="local_icon" size="1.5rem"/>
					</van-col>
					<van-col span="18" class="child2">
						<van-row>
							<van-col span="12">{{currentAddr.ReceiverName}}</van-col>
							<van-col span="12">{{currentAddr.ReceiverMobile}}</van-col>
						</van-row>
						<van-row>
							<van-col span="24">{{currentAddr.ReceiverProvince+currentAddr.ReceiverCity+currentAddr.ReceiverDistrict+currentAddr.DetailAddress}}</van-col>
						</van-row>
					</van-col>
					<van-col span="2" class="child3">
						<van-icon name="arrow" />
					</van-col>
				</van-row>
				<div class="line"></div>
			</div>

			<div v-if="fromPage === 'goodsDetail'">
				<van-card
					:num="data.goodsNum"
					:price="data.Price"
					:title="data.Name"
					:thumb="data.MainImage"
					currency="Block"
				/>
				<van-row class="quantity">
					<van-col span="12">购买数量</van-col>
					<van-col span="12">
						<van-stepper v-model="data.goodsNum" :max="data.Stock" disable-input integer/>
					</van-col>
				</van-row>
				<van-cell-group class="delivery" >
					<van-cell title="配送方式" value="快递 免邮" is-link/>
				</van-cell-group>
				<van-cell-group class="totalPrice">
					<van-cell title="商品金额" :value="`Block ${totalPrice.toFixed(2)}`" />
				</van-cell-group>
			</div>
			<div v-else>
				<div v-for="(item,index) in data" :key="index">
					<van-card
						:num="item.Quantity"
						:price="item.Price"
						:title="item.ProductName"
						:thumb="item.MainImage"
						currency="Block"
					/>
					<van-cell-group class="delivery" >
						<van-cell title="配送方式" value="快递 免邮" is-link/>
					</van-cell-group>
					<van-cell-group class="totalPrice">
						<van-cell title="商品金额" :value="`Block ${singleTotalPrice(item.Quantity,item.Price)}`" />
					</van-cell-group>
				</div>
			</div>
		</div>
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show"></van-loading>
		<van-submit-bar 
			:price="totalPrice*100"
			button-text="提交订单"
			currency="Block"
			safe-area-inset-bottom
			@submit="onSubmit"
		/>
		
	</div>
</template>

<script>
	import {NavBar,Row,Col,CellGroup,Cell,Button,Image,SubmitBar,Card,Icon,Stepper,Overlay,Loading,Toast,ContactCard} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Row.name]:Row,
			[Col.name]:Col,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[Button.name]:Button,
			[Image.name]:Image,
			[SubmitBar.name]:SubmitBar,
			[Card.name]:Card,
			[Icon.name]:Icon,
			[Stepper.name]:Stepper,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading,
			[Toast.name]:Toast,
			[ContactCard.name]:ContactCard
		},
		data(){
			return{
				fromPage:'',
				data:'',
				currentAddr:'',
				local_icon:Images.shop.ic_address_icon,
				line:Images.shop.ic_line,
				show:false
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			getData(){
				if (this.fromPage === 'goodsDetail') {
					this.data = this.$store.state.confirmOrder
				}
				if (this.fromPage === 'cart') {
					this.data = []
					for (var i = 0; i < this.$store.state.allDatas.data.length; i++) {
						if (this.$store.state.allDatas.data[i].Checked) {
							this.data.push(this.$store.state.allDatas.data[i])
						}
					}
				}
				this.currentAddr = this.$store.state.currentReceiverAddress
			},
			selectAddr(){
				this.$router.push('/addressSetting')
			},
			onSubmit(){
				if (this.fromPage === 'goodsDetail') {
					this.show = true
					var data = {
						UserID:this.$store.state.userID,
						ProductID:this.data.goodsID,
						Quantity:this.data.goodsNum,
						ShippingID:this.$store.state.currentReceiverAddress.ID
					}
					var url = '/order/buyDirect?token='+this.$store.state.token
					this.axios.post(url,data).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								this.show = false
								this.$store.commit('paymentData',response.data.obj)
								this.$router.push('/paymentPage')
							}, 1000)
						} 
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}
				if (this.fromPage === 'cart') {
					this.show = true
					var url = '/order/createOrder?user_id='+this.$store.state.userID+'&ship_id='+this.$store.state.currentReceiverAddress.ID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if(response.data.status === 1){
							setTimeout(() => {
								this.$store.commit('deleteGoodFromCart')
								this.show = false
								this.$store.commit('paymentData',response.data.obj)
								this.$router.push('/paymentPage')
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}
			}
		},
		watch:{
			goodsNum(){
				this.data.goodsNum = this.data.goodsNum <= this.data.Stock ? this.data.goodsNum : this.data.Stock
			}
		},
		computed:{
			singleTotalPrice(){
				return function (num,price) {
					return (num*price).toFixed(2)
				}
			},
			totalPrice(){
				if (this.fromPage === 'goodsDetail') {
					return this.data.Price*this.data.goodsNum
				}
				if (this.fromPage === 'cart') {
					var totalPrice = 0
					for (var i = 0; i < this.data.length; i++) {
						totalPrice += this.data[i].Quantity * this.data[i].Price
					}
					return totalPrice
				}
			}
		},
		beforeRouteEnter(to,from,next){
			next(vm=>{
				vm.fromPage = from.name
				vm.getData()
			})
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: #000000;
	}
	.content{
		margin: 2.875rem auto 3.125rem;
	}
	.address{
		background: white;
	}
	.address .child1 .van-icon,.address .child3 .van-icon{
		margin: 1.05rem auto;
	}
	.address .child2{
		text-align: left;
	}
	.address .child2 .van-row{
		margin: 0.5rem auto;
		font-size: 0.8125rem;
	}
	.line{
		background: url(../assets/Images/Shop/orderLine.png);
		background-size: 100%;
		width: 100%;
		height: 0.3125rem;
	}
	.van-card{
		background: white;
		margin-top: 0.3125rem;
	}
	.van-card__title{
		text-align: left;
	}
	.van-card__price{
		float: left;
	}
	.quantity{
		background: white;
		border-top: 0.0625rem solid #F0F0F0;
	}
	.quantity .van-col{
		padding: 0.625rem 1rem;
		font-size: 0.875rem;
		line-height: 1.5rem;
	}
	.quantity .van-col:first-child{
		text-align: left;
		font-size: 0.9375rem;
	}
	.quantity .van-col:last-child{
		text-align: right;
	}
	.delivery{
		text-align: left;
	}
	.totalPrice{
		text-align: left;
	}
	.totalPrice .van-cell__value{
		color: red;
	}
	.delivery .van-cell,.totalPrice .van-cell{
		font-size: 0.875rem;
		color: #999999;
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
