<template>
	<div class="orderDetail">
		<van-nav-bar
			title="订单详情"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		<div class="content">
			<div>
				<van-row class="title">
					<van-col span="16">
						<van-row class="title1" >
							<van-col span="24">{{title1}}</van-col>
						</van-row>
						<van-row class="title2">
							<van-col span="24">{{title2}}</van-col>
						</van-row>
					</van-col>
					<van-col span="8">
						<van-image :src="wallet" width="3.75rem" height="3.75rem"/>
					</van-col>
				</van-row>
				<van-row class="address">
					<van-col span="4" class="child1">
						<van-icon :name="local_icon" size="2rem"/>
					</van-col>
					<van-col span="20" class="child2">
						<van-row>
							<van-col span="12">{{shippingAddress.ReceiverName}}</van-col>
							<van-col span="12">{{shippingAddress.ReceiverMobile}}</van-col>
						</van-row>
						<van-row>
							<van-col span="24">{{shippingAddress.ReceiverProvince+shippingAddress.ReceiverCity+shippingAddress.ReceiverDistrict+shippingAddress.DetailAddress}}</van-col>
						</van-row>
					</van-col>
				</van-row>
				<div class="line"></div>
				<van-row v-if="title1 === '商品正在路上' || title1 === '交易已经完成'" class="logistics" >
					<van-col span="4" class="child1">
						<van-icon :name="logistics_icon" size="2rem"/>
					</van-col>
					<van-col span="20" class="child2">
						<van-row>
							<van-col span="24">快递公司:{{data.ShippingCompany}}</van-col>
						</van-row>
						<van-row>
							<van-col span="24">快递单号:{{data.ShippingNum}}</van-col>
						</van-row>
						<van-row>
							<van-col span="24">发货时间:{{data.ShippingTime.replace('T',' ').replace('Z',' ')}}</van-col>
						</van-row>
					</van-col>
				</van-row>
			</div>
		
			<div>
				<van-card
					:num="data.Quantity"
					:price="data.CurrentUnitPrice"
					:title="data.ProductName"
					:thumb="data.ProductImage"
					currency="Block"
					@click="getGoodsDetailByID(data.ProductID)"
				/>
				<van-cell-group class="delivery" >
					<van-cell title="配送方式" value="快递 免邮" is-link/>
				</van-cell-group>
				<van-cell-group class="totalPrice">
					<van-cell title="商品金额" :value="`Block ${(data.TotalPrice).toFixed(2)}`" />
				</van-cell-group>
			</div>
			
			<div v-if="data.title ? data.title === '退款' : false">
				<van-row class="refund">
					<van-col span="24">
						<van-row class="info1">
							<van-col span="24">退款原因</van-col>
						</van-row>
						<van-row class="info2">
							<van-col>{{data.Reason}}</van-col>
						</van-row>
					</van-col>
				</van-row>
			</div>
			
			<div v-if="data.title ? data.title === '退货' : false">
				<van-row class="refund">
					<van-col span="24">
						<van-row class="info1">
							<van-col span="24">退货原因</van-col>
						</van-row>
						<van-row class="info2">
							<van-col>{{data.Reason}}</van-col>
						</van-row>
					</van-col>
				</van-row>
			</div>
			
			<div v-if="title1 === '退款失败' || title1 === '退货申请失败' || title1 === '卖家拒绝退款'">
				<van-row class="refund">
					<van-col span="24">
						<van-row class="info1">
							<van-col span="24">拒绝原因</van-col>
						</van-row>
						<van-row class="info2">
							<van-col>{{data.Reject}}</van-col>
						</van-row>
					</van-col>
				</van-row>
			</div>
			
			<div v-if="title1 === '等待卖家收货'">
				<van-row class="refund">
					<van-col span="24">
						<van-row class="info1">
							<van-col span="24">退货接收人信息</van-col>
						</van-row>
						<van-row class="info2">
							<van-col>收货人姓名:{{data.ShipName}}</van-col>
						</van-row>
						<van-row class="info2">
							<van-col>收货人电话:{{data.ShipPhone}}</van-col>
						</van-row>
						<van-row class="info2">
							<van-col>收货地址:{{data.ShipAddress}}</van-col>
						</van-row>
					</van-col>
				</van-row>
				<van-row class="refund">
					<van-col span="24">
						<van-row class="info1">
							<van-col span="24">快递单号</van-col>
						</van-row>
						<van-row>
							<van-col span="24">
								<van-field v-model="orderNo" placeholder="请填写退货商品的快递单号"/>
							</van-col>
						</van-row>
					</van-col>
				</van-row>
			</div>
			
			<div>
				<van-row class="orderInfo">
					<van-col span="24">
						<van-row class="info1">
							<van-col span="24">订单信息</van-col>
						</van-row>
						<van-row class="info2">
							<van-col span="5">订单编号：</van-col>
							<van-col span="19">{{data.OrderNo}}</van-col>
						</van-row>
						<van-row class="info3">
							<van-col span="5">创建时间：</van-col>
							<van-col span="19">{{data.CreateTime.replace('T',' ').replace('Z',' ')}}</van-col>
						</van-row>
					</van-col>
				</van-row>
			</div>
		</div>
		
		<van-row v-if="title1 === '等待买家付款'" class="bottomBtn">
			<van-col span="6" offset="12">
				<van-button type="warning" round @click="cancelOrder">取消订单</van-button>
			</van-col>
			<van-col span="6">
				<van-button type="danger" round @click="pay">付款</van-button>
			</van-col>
		</van-row>
		<van-row v-if="title1 === '等待卖家发货'" class="bottomBtn">
			<van-col span="6" offset="12">
				<van-button type="warning" round @click="refund">退款</van-button>
			</van-col>
			<van-col span="6">
				<van-button type="danger" round @click="confirmReceipt">确认收货</van-button>
			</van-col>
		</van-row>
		<van-row v-if="title1 === '商品正在路上'" class="bottomBtn">
			<van-col span="6" offset="6">
				<van-button type="default" round @click="logistics">查看物流</van-button>
			</van-col>
			<van-col span="6">
				<van-button type="warning" round @click="returnOrder">退货</van-button>
			</van-col>
			<van-col span="6">
				<van-button type="danger" round @click="confirmReceipt">确认收货</van-button>
			</van-col>
		</van-row>
		<van-row v-if="title1 === '交易已经完成' || title1 === '交易已经取消'" class="bottomBtn">
			<van-col span="6" offset="18">
				<van-button type="danger" round @click="cancelOrder">删除订单</van-button>
			</van-col>
		</van-row>
		<van-row v-if="title1 === '退款审核中'" class="bottomBtn">
			<van-col span="6" offset="18">
				<van-button type="danger" round @click="cancelRefund">取消退款</van-button>
			</van-col>
		</van-row>
		<van-row v-if="title1 === '退款失败'" class="bottomBtn">
			<van-col span="6" offset="12">
				<van-button type="warning" round @click="cancelRefund">取消退款</van-button>
			</van-col>
			<van-col span="6">
				<van-button type="danger" round @click="confirmReceipt">确认收货</van-button>
			</van-col>
		</van-row>
		<van-row v-if="title1 === '等待卖家收货'" class="bottomBtn">
			<van-col span="6" offset="18">
				<van-button type="danger" round @click="submitOrderNo">提交</van-button>
			</van-col>
		</van-row>
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show"></van-loading>
	</div>
</template>

<script>
	import {NavBar,Row,Col,CellGroup,Cell,Button,Image,Card,Icon,Overlay,Loading,Toast,Dialog,Field} from 'vant'
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
			[Card.name]:Card,
			[Icon.name]:Icon,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading,
			[Toast.name]:Toast,
			[Dialog.Component.name]:Dialog.Component,
			[Field.name]:Field
		},
		data(){
			return{
				data:'',
				title1:'',
				title2:'',
				show:false,
				local_icon:Images.indent.ic_addr,
				wallet:Images.indent.ic_wallet,
				logistics_icon:Images.indent.ic_wuli,
				shippingAddress:'',
				orderNo:''
			}
		},
		methods:{
			goBack(){
				if (this.data.title === '退款') {
					this.$router.push('/refundAndReturnPage/退款')
				} else if (this.data.title === '退货') {
					this.$router.push('/refundAndReturnPage/退货')
				}else{
					this.$router.back()
				}
			},
			getTitle(){
				if (this.data.Status === 0) {
					this.title1 = '交易已经取消'
					this.title2 = '赶快重新下单吧'
				}
				if (this.data.Status === 10) {
					this.title1 = '等待买家付款'
					this.title2 = '请您尽快付款哟'
				}
				if (this.data.Status === 20) {
					this.title1 = '等待卖家发货'
					this.title2 = '商品正在出仓哟'
				}
				if (this.data.Status === 30) {
					this.title1 = '商品正在路上'
					this.title2 = '请您耐心等待哟'
				}
				if (this.data.Status === 40) {
					this.title1 = '交易已经完成'
					this.title2 = '快去再下一单吧'
				}
				if (this.data.Status === 50) {
					this.title1 = '交易已经关闭'
					this.title2 = ''
				}
				if (this.data.Status === 21 || (this.data.title === '退款' && this.data.Status === 1)) {
					this.title1 = '退款审核中'
					this.title2 = '请您耐心等待'
				}
				if (this.data.Status === 22 || (this.data.title === '退款' && this.data.Status === 2)) {
					this.title1 = '退款成功'
					this.title2 = '退款会尽快退回到您的账户'
				}
				if (this.data.Status === 23 || (this.data.title === '退款' && this.data.Status === 3)) {
					this.title1 = '退款失败'
					this.title2 = '您的退款申请未通过'
				}
				if (this.data.title === '退货' && this.data.Status === 1) {
					this.title1 = '退货审核中'
					this.title2 = '请您耐心等待'
				}
				if (this.data.title === '退货' && this.data.Status === 2) {
					this.title1 = '退货申请成功'
					this.title2 = '等待卖家上传收货信息'
				}
				if (this.data.title === '退货' && this.data.Status === 3) {
					this.title1 = '退货申请失败'
					this.title2 = '您的退货申请未通过'
				}
				if (this.data.title === '退货' && this.data.Status === 4) {
					this.title1 = '等待卖家收货'
					this.title2 = '请您提交快递单号后耐心等待'
				}
				if (this.data.title === '退货' && this.data.Status === 5) {
					this.title1 = '卖家同意退款'
					this.title2 = '退款会尽快退回到您的账户'
				}
				if (this.data.title === '退货' && this.data.Status === 6) {
					this.title1 = '卖家拒绝退款'
					this.title2 = '您购买的商品不能进行退款'
				}
			},
			getShippingAddress(){
				var url = '/order/getshipByOrder?order_no='+this.data.OrderNo+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						this.shippingAddress = response.data.obj
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			cancelOrder(){
				Dialog.confirm({
					title:'提示',
					message:'是否确认取消订单'
				}).then(()=>{
					this.show = true
					var url = '/order/cancelOrder?order_no='+this.data.OrderNo+'&user_id='+this.data.UserID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast('订单已取消')
								this.show = false
								this.$router.back()
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			},
			pay(){
				this.$store.commit('paymentData',this.data)
				this.$router.push('/paymentPage')
			},
			getGoodsDetailByID(pid){
				var url = '/product/getProductByID?pid='+pid+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						this.$store.commit("goodsDetail",response.data.obj)
						this.$router.push('/goodsDetail')
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			refund(){
				this.$store.commit('refundReturnData',this.data)
				this.$router.push('/refundApply/退款申请')
			},
			confirmReceipt(){
				Dialog.confirm({
					title:'提示',
					message:'是否确认收货'
				}).then(()=>{
					this.show = true
					var url = '/order/ensureOrder?order_no='+this.data.OrderNo+'&user_id='+this.data.UserID+'&product_id='+this.data.ProductID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast('已确认收货')
								this.show = false
								this.$router.back()
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			},
			logistics(){
				
			},
			returnOrder(){
				this.$store.commit('refundReturnData',this.data)
				this.$router.push('/returnApply/退货申请')
			},
			cancelRefund(){
				Dialog.confirm({
					title:'提示',
					message:'是否取消退款'
				}).then(()=>{
					this.show = true
					var url = '/order/cancelReturn?user_id='+this.$store.state.userID+'&order_no='+this.data.OrderNo+'&product_id='+this.data.ProductID+'&return_id='+this.data.ID+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(() => {
								Toast(response.data.detail)
								this.show = false
								this.$router.back()
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}).catch(()=>{})
			},
			submitOrderNo(){
				if (!this.orderNo) {
					Toast('请填写快递单号')
					return
				}
				var url = '/order/addReturnLogistic?rid='+this.data.ID+'&logistic_no='+this.orderNo+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						setTimeout(() => {
							Toast(response.data.detail)
							this.show = false
							this.$router.back()
						}, 1000);
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			}
		},
		beforeMount() {
			this.data = this.$store.state.orderDetail
			console.log(this.data)
			this.getTitle()
			this.getShippingAddress()
		}
	}
</script>

<style scoped>
	.van-nav-bar{
		background: rgb(255, 115, 105);
	}
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: white;
	}
	.van-nav-bar__title{
		color: white;
	}
	.content{
		margin: 2.875rem auto 3.125rem;
	}
	.title{
		background: rgb(255, 115, 105);
		color: white;
	}
	.title1 {
		font-size: 1.0625rem;
		margin: 0.625rem 1.25rem;
		text-align: left;
	}
	.title2{
		font-size: 0.875rem;
		margin: 0.625rem 1.25rem;
		text-align: left;
	}
	.title .van-image{
		margin: 0.625rem auto 0;
	}
	.address{
		background: white;
	}
	.address .child1 .van-icon{
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
		background: url(../../assets/Images/Shop/orderLine.png);
		background-size: 100%;
		width: 100%;
		height: 0.3125rem;
	}
	.logistics{
		background: white;
	}
	.logistics .child1 .van-icon{
		margin: 1.4rem auto;
	}
	.logistics .child2{
		text-align: left;
		color: #999999;
		font-size: 0.8125rem;
	}
	.logistics .child2 .van-row{
		margin: 0.3125rem;
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
	.delivery .van-cell,.totalPrice .van-cell{
		font-size: 0.875rem;
		color: #999999;
	}
	.totalPrice{
		text-align: left;
	}
	.totalPrice .van-cell__value{
		color: red;
	}
	.refund{
		text-align: left;
		background: white;
		margin-top: 0.3125rem;
	}
	.refund .info1{
		font-size: 0.875rem;
		color: #808080;
		margin: 0.625rem 0.9375rem;
	}
	.refund .info2{
		font-size: 0.75rem;
		color: #999999;
		margin: 0.625rem 0.9375rem;
	}
	.refund .van-field{
		border: 0.0625rem #F0F0F0 solid;
	}
	.orderInfo{
		text-align: left;	
		background: white;
		margin-top: 0.3125rem;
	}
	.orderInfo .info1{
		font-size: 0.875rem;
		color: #808080;
		margin: 0.625rem 0.9375rem;
	}
	.orderInfo .info2,.orderInfo .info3{
		font-size: 0.75rem;
		color: #999999;
		margin: 0.625rem 0.9375rem;
	}
	.bottomBtn{
		width: 100%;
		height: 3.125rem;
		line-height: 3.125rem;
		position: fixed;
		bottom: 0;
		background: white;
	}
	.bottomBtn .van-button{
		height: 2.2rem;
		line-height: 2.2rem;
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
