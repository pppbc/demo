<template>
	<div id="detail">
		
		<van-nav-bar v-show="!showTab" :border="false" @click-left="goBack" fixed>
			<van-icon class="backBtn" :name="backImg" slot="left" />
		</van-nav-bar>
		
		<transition name="fade" mode="out-in">
			<van-row v-show="showTab"  class="hidetab">
				<van-col span="4">
					<van-icon name="arrow-left" size="1.5rem" @click="goBack"/>
				</van-col>
				<van-col span="16">
					<van-button @click="goPos(1)">商品</van-button>
					<van-button @click="goPos(2)">评价</van-button>
					<van-button @click="goPos(3)">详情</van-button>
				</van-col>
			</van-row>
		</transition>
		
		<div class="container" ref="pos1">
			<van-row>
				<van-col span="24">
					<img class="mainImg" :src="goodsDetail.MainImage">
				</van-col>
			</van-row>
			<van-row class="price">
				<van-col span="24"><span>{{goodsDetail.Price}}</span>Block</van-col>
			</van-row>
			<van-row class="title">
				<van-col span="24">{{goodsDetail.Name}}</van-col>
			</van-row>
			<van-row class="desc">
				<van-col span="24">{{goodsDetail.Description}}</van-col>
			</van-row>
			<van-row class="icon">
				<van-col span="24">
					<van-icon :name="icon"/><span>全场包邮</span>
					<van-icon :name="icon"/><span>7天退换</span>
					<van-icon :name="icon"/><span>假一赔十</span>
				</van-col>
			</van-row>
			<van-cell-group class="inventory">
				<van-cell title="商品剩余数量:" :value="goodsDetail.Stock" value-class="stockValue"/>
			</van-cell-group>
			<van-row class="quantity">
				<van-col span="12">购买数量：</van-col>
				<van-col span="12">
					<van-stepper v-model="goodsNum" :max="goodsDetail.Stock" integer/>
				</van-col>
			</van-row>
			<van-cell-group class="totalPrice">
				<van-cell title="总价格:" :value="totalPrice" value-class="stockValue"/>
			</van-cell-group>
			<van-cell-group class="commits" ref="pos2">
				<van-cell :title="`商品评价(${comments})`" value="查看全部" is-link @click="goComments"/>
				<van-cell class="com1" v-if="allComments.length > 0" :title="username" :value="allComments[0].Title" :label="allComments[0].Content"/>
				<van-cell class="com2" v-else title="该商品还没有相关评论" />
			</van-cell-group>
			<span id="gd" ref="pos3">———— 商品详情 ————</span>
			<van-row class="subImg">
				<van-col span="24" v-for="(item,index) in subImages" :key="index">
					<van-image
						:src="item.PicUrl"
						lazy-load
						width="100%"
						height="auto"
					/>
				</van-col>
			</van-row>
			
			<span id="gd">———— 已经到底了 ————</span>
			<van-icon v-if="showTab" class="upArrow" :name="upArrow" size="2.5rem" @click="scrollTop"/>
		</div>
		
		
	
		<van-goods-action>
			<van-goods-action-icon icon="chat-o" text="客服" @click="onClickChat"/>
			<van-goods-action-icon icon="cart-o" text="购物车" :info="cartInfo" @click="onCartClick" />
			<van-goods-action-button type="warning" text="加入购物车" @click="onAddToCart"/>
			<van-goods-action-button type="danger" text="立即购买" @click="onBuyClick"/>
		</van-goods-action>
		
	</div>
</template>

<script>
	import {NavBar,Icon,Row,Col,Button,Image,CellGroup,Cell,GoodsAction,GoodsActionIcon,GoodsActionButton,Stepper,Toast} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Icon.name]:Icon,
			[Row.name]:Row,
			[Col.name]:Col,
			[Button.name]:Button,
			[Image.name]:Image,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[GoodsAction.name]:GoodsAction,
			[GoodsActionIcon.name]:GoodsActionIcon,
			[GoodsActionButton.name]:GoodsActionButton,
			[Stepper.name]:Stepper,
			[Toast.name]:Toast
		},
		data(){
			return{
				backImg:Images.shop.ic_back,
				goodsDetail:'',
				showTab:false,
				icon:Images.shop.ic_item,
				goodsNum:1,
				comments:0,
				allComments:[],
				upArrow:Images.common.ic_up_arrow,
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive',this.$store.state.fromPage)
				this.$router.back()
			},
			getAllComments(count){
				var url = '/order/getProductComment?product_id='+this.goodsDetail.ID+'&page='+count
				this.axios.get(url).then((response)=>{
					if (response.data.status === 1) {
						this.$store.commit("allComments",response.data.obj)
						this.allComments = response.data.obj
					} 
				}).catch((response)=>{
					Toast("网络请求出错")
				})
			},
			onCartClick(){
				this.$router.push('/cart')
			},
			onAddToCart(){
				var url = '/cart/add?user_id='+this.$store.state.userID+'&product_id='+this.goodsDetail.ID+'&quantity='+this.goodsNum+'&token='+this.$store.state.token
				// var data = {
				// 	user_id:this.$store.state.userID,
				// 	ProductID:this.goodsDetail.ID,
				// 	Quantity:this.goodsNum,
				// 	Checked:true
				// }
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						this.updateCartScreen()
						Toast("添加成功")
					} else{
						Toast("添加失败，请重试")
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			onBuyClick(){
				var data = {
					MainImage:this.goodsDetail.MainImage,
					Name:this.goodsDetail.Name,
					goodsNum:this.goodsNum,
					Price:this.goodsDetail.Price,
					goodsID:this.goodsDetail.ID,
					Stock:this.goodsDetail.Stock
				}
				this.$store.commit('confirmOrder',data)
				this.$router.push('/confirmOrder')
			},
			handleScroll(){
				var scrollY = document.documentElement.scrollTop || document.body.scrollTop
				if (scrollY > 200) {
					this.showTab = true
				} else{
					this.showTab = false
				}
			},
			scrollTop(){
				var top = document.documentElement.scrollTop || document.body.scrollTop
				const timeTop = setInterval(()=>{
					document.body.scrollTop = document.documentElement.scrollTop = top -= 50
					if (top <= 0) {
						clearInterval(timeTop)
					}
				},10)
			},
			updateCartScreen(){
				var goodsID = this.goodsDetail.ID
				var index = 0
				if (this.$store.state.allDatas.data) {
					index = this.$store.state.allDatas.data.findIndex((item)=>{
						return item.ProductID === goodsID
					})
				}
				if (index === -1) {
					this.$store.commit({type:'addToCart1',data:this.goodsDetail,num:this.goodsNum})
				}else{
					this.$store.commit({type:'addToCart2',index:index,num:this.goodsNum})
				}
			},
			goComments(){
				this.$router.push('/allComments/'+this.goodsDetail.ID)
			},
			goPos(index){
				switch (index){
					case 1:
						window.scrollTo(0,this.$refs.pos1.offsetTop)
						break;
					case 2:
						window.scrollTo(0,this.$refs.pos2.offsetTop - 25)
						break;
					case 3:
						window.scrollTo(0,this.$refs.pos3.offsetTop - 25)
						break
					default:
						break;
				}
			},
			onClickChat(){
				console.log('通过商家客服id进入聊天界面')
			}
		},
		computed:{
			totalPrice(){
				var total = this.goodsDetail.Price * this.goodsNum
				total = parseFloat(total)
				total = total.toFixed(2)
				total = total.toLocaleString()
				return total+" Block"
			},
			username(){
				return this.allComments[0].UserName.replace(this.allComments[0].UserName.substring(3,7),"****")
			},
			subImages(){
				return this.goodsDetail.SubImages ? JSON.parse(this.goodsDetail.SubImages) : []
			},
			cartInfo(){
				return this.$store.state.allDatas.data.length
			}
		},
		watch:{
			goodsNum(){
				this.goodsNum = this.goodsNum <= this.goodsDetail.Stock ? this.goodsNum : this.goodsDetail.Stock
			}
		},
		beforeMount() {
			this.goodsDetail = this.$store.state.goodsDetail
			this.$store.commit("requestAllCommentsCount",0)
			this.getAllComments(this.$store.state.requestAllCommentsCount)
		},
		mounted(){
			window.addEventListener('scroll',this.handleScroll)
		}
	}
</script>

<style scoped>
	#detail{
		background: white;
	}
	.van-nav-bar{
		background: rgba(0,0,0,0);
		z-index: 1;
	}
	.backBtn{
		font-size: 1.5rem;
	}
	.hidetab{
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		height: 2rem;
		line-height: 2rem;
		z-index: 1;
		background: white;
	}
	.hidetab .van-col:first-child{
		text-align: left;
		padding-left: 0.3125rem;
		padding-top: 0.25rem;
	}
	.hidetab .van-button{
		height: 2rem;
		line-height: 2rem;
		border: none;
		background: rgba(0,0,0,0);
		color: gray;
	}
	.container{
		margin-bottom: 3.125rem;
	}
	.mainImg{
		width: 100%;
		height: auto;
		max-height: 25rem;
	}
	.price{
		color: red;
		font-size: 0.75rem;
		text-align: left;
	}
	.price span{
		font-size: 1.25rem;
		margin-left: 0.625rem;
	}
	.title{
		font-size: 0.9375rem;
		text-align: left;
		margin: 0.3125rem 0.625rem 0;
		font-weight: bold;
	}
	.desc{
		font-size: 0.75rem;
		text-align: left;
		margin: 0.3125rem 0.625rem 0;
		color: #564F5F;
	}
	.icon{
		text-align: left;
		margin: 0.3125rem 0.625rem 0;
	}
	.icon .van-icon{
		font-size: 0.625rem;
	}
	.icon span{
		margin-left: 0.3125rem;
		margin-right: 0.625rem;
		font-size: 0.75rem;
	}
	.van-cell:after{
		border: none;
	}
	.stockValue{
		font-size: 1.25rem;
		color: red;
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
	.inventory,.totalPrice,.commits{
		text-align: left;
	}
	.commits .van-cell__value:last-child{
		text-align: left;
	}
	.commits .com2 {
		font-size: 0.875rem;
		text-align: center;
		color: #999999;
	}
	#gd{
		text-align: center;
		display: block;
		height: 1.875rem;
		line-height: 1.875rem;
		font-size: 0.625rem;
		color: gray;
		background: #F5F5F5;
	}
	.subImg .van-col,.subImg .van-image{
		border: 0 none;
		margin: 0 0 -3px;
		padding: 0;
	}
	.upArrow{
		position: fixed;
		right: 0.9375rem;
		bottom: 4.375rem;
	}
</style>

