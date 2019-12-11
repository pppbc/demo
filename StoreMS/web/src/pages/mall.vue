<template>
	<div id="mall">
		<img class="bg" :src="bg" alt="">
		<van-nav-bar 
			title="商城"
			:border="border"
			fixed
		>
			<van-icon name="shopping-cart-o" slot="right" @click="goCart"/>
		</van-nav-bar>
		
		<div v-if="isonline">
			<my-integral :avatar="avatar" :allBlock="allBlock"></my-integral>
			
			<swipe :swipeItems="allActicities"></swipe>
			
			<van-row>
				<van-col span="6">
					<van-button class="btn1">
						<div class="t1">现在</div>
						<div class="t2">可兑换商品</div>
					</van-button>
				</van-col>
				<van-col span="6">
					<van-button class="btn2">
						<div class="t1">1个月后</div>
						<div class="t2">可兑换商品</div>
					</van-button>
				</van-col>
				<van-col span="6">
					<van-button class="btn3">
						<div class="t1">3个月后</div>
						<div class="t2">可兑换商品</div>
					</van-button>
				</van-col>
				<van-col span="6">
					<van-button class="btn4">
						<div class="t1">6个月后</div>
						<div class="t2">可兑换商品</div>
					</van-button>
				</van-col>
			</van-row>
				
			<goods-tabs></goods-tabs>
		</div>
		<div v-else class="offline">
			<div>
				<van-image :src="noNetwork" />
				<van-button plain type="info" @click="onRefresh">点击刷新</van-button>
			</div>
		</div>
		
		
	</div>
</template>

<script>
	import {NavBar,Toast,Button,Row,Col,Icon} from 'vant'
	import {Images} from '../resources/index.js'
	import myIntegral from '@/components/myIntegral.vue'
	import swipe from '@/components/swipe.vue'
	import goodsTabs from '@/components/goodsTabs.vue'
	export default{
		inject:['reload'],
		components:{
			[NavBar.name]:NavBar,
			[Toast.name]:Toast,
			[Button.name]:Button,
			[Row.name]:Row,
			[Col.name]:Col,
			[Icon.name]:Icon,
			myIntegral,
			swipe,
			goodsTabs
		},
		data(){
			return{
				bg:Images.common.ic_header_bg,
				border:false,
				goods:[],
				category:[],
				categoryIcon : Images.shop.ic_category_icon,
				categoryAlls : [{"ID":0,"ParentID":0,"Name":"全部商品","Status":0,"SortOrder":0,"Detail":"","CreateTime":"","UpdateTime":""}],
				isonline:navigator.onLine,
				noNetwork:Images.common.ic_no_network
			}
		},
		methods:{
			getCategory(){
				var url = '/product/getCat'
				this.axios.get(url).then((response)=>{
					if (response.data.obj.length > 0) {
						this.category = [...this.category,...response.data.obj]
						this.categoryAll()
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			categoryAll(){
				this.category.map((item,index)=>{
					this.categoryAlls.push(item)
				})
			},
			netApi(count,pid){
				if (pid === 0) {
					var url = '/product/allList?page='+count
					this.axios.get(url).then((response)=>{
						if (response.data.status === 1 && response.data.obj.length > 0) {
							this.$store.commit('requestMallcount2')
							this.goods = [...this.goods,...response.data.obj]
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				} else{
					var url = '/product/getProductList?page='+count+'&category='+pid
					this.axios.get(url).then((response)=>{
						this.$store.commit('requestMallcount2')
						this.goods = [...this.goods,...response.data.obj]
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}
			},
			updateOnlineStatus(e){
				const {type} = e
				this.isonline = type === 'online'
			},
			onRefresh(){
				this.$store.commit('changeTabbarActive','mall')
				this.$nextTick(()=>{
					this.reload()
				})
			},
			goCart(){
				this.$store.commit('fromPage','mall')
				this.$router.push('/cart')
			}
		},
		computed:{
			avatar(){
				return this.$store.state.avatar
			},
			allActicities(){
				return this.$store.state.allActicities.slice(0,5)
			},
			allBlock(){
				return this.$store.getters.allBlock
			}
		},
		beforeMount() {
			this.getCategory()
			this.netApi(this.$store.state.requestMallcount1,0)
		},
		mounted() {
			window.addEventListener('online',this.updateOnlineStatus)
			window.addEventListener('offline',this.updateOnlineStatus)
		},
		beforeDestroy() {
			window.removeEventListener('online',this.updateOnlineStatus)
			window.removeEventListener('offline',this.updateOnlineStatus)
		},
		watch:{
			isonline(){
				this.isonline = navigator.onLine
			}
		}
	}
</script>

<style scoped>
	.bg{
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 6.5rem;
		z-index: -1;
	}
	.van-nav-bar{
		background: url(../assets/Images/Home/header_bg.png);
	}
	.van-nav-bar__title{
		color: #fff;
		font-size: 1.125rem;
	}
	.van-nav-bar .van-icon{
		color: #FFFFFF;
		font-size: 1.5rem;
	}
	.van-row{
		margin: 0.625rem 0.5rem 0;
	}
	
	.van-row .van-button{
		color: white;
		height: 2.75rem;
		line-height: 1.375rem;
		padding: 0;
		width: 100%;
	}
	.btn1{
		background: url(../assets/Images/Shop/time_1.png) no-repeat center;
		background-size: 100% 100%;
	}
	.btn2{
		background: url(../assets/Images/Shop/time_2.png) no-repeat center;
		background-size: 100% 100%;
	}
	.btn3{
		background: url(../assets/Images/Shop/time_3.png) no-repeat center;
		background-size: 100% 100%;
	}
	.btn4{
		background: url(../assets/Images/Shop/time_4.png) no-repeat center;
		background-size: 100% 100%;
	}
	.t1{
		margin: 0;
		padding: 0;
		font-size: 0.9375rem;
	}
	.t2{
		margin: 0;
		padding: 0;
		font-size: 0.625rem;
	}
	.offline{
		width: 100%;
		margin: 0 auto;
	}
	.offline div{
		width: 50%;
		position: absolute;
		left: 50%;
		top: 20%;
	}
	.offline .van-image{
		width: 100%;
		height: auto;
		position: relative;
		left: -50%;
	}
	.offline .van-button{
		width: 6.25rem;
		position: relative;
		left: -50%;
		margin-top: 0.625rem;
	}
</style>
