<template>
	<div id="wallet">
		<img :src="bg" class="bg">
		<van-nav-bar
			title="我的钱包"
			left-arrow
			fixed
			:border="false"
			@click-left="goBack"
		/>
		
		<van-row class="wallet">
			<van-col class="wallet01" span="24">
				<van-image :src="walletImg" width="3.125rem" height="3.125rem"/>
			</van-col>
			<van-col class="wallet02" span="24">余额：{{$store.state.totalBlock}}积分</van-col>
		</van-row>
			
		<div class="list">
			<van-row v-for="(item,index) in integralData" :key="index">
					<van-col class="list01" span="5">
						<van-image :src="item.PointLogo" width="3.75rem" height="3.75rem"/>
					</van-col>
					<van-col class="list02" span="19">
						<van-row>
							<van-col span="24">{{item.PointName}}</van-col>
						</van-row>
						<van-row>
							<van-col span="12">兑换比例{{item.PointFate}}</van-col>
							<van-col span="12">余额：{{item.Balance}}</van-col>
						</van-row>
						<van-row>
							<van-col span="24">过期时间{{item.ExpireTime.replace("T"," ").replace("Z"," ")}}</van-col>
						</van-row>
					</van-col>
			</van-row>
		</div>
		
	</div>
</template>

<script>
	import {NavBar,Row,Col,Image} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Row.name]:Row,
			[Col.name]:Col,
			[Image.name]:Image
		},
		data(){
			return {
				bg:Images.mineView.ic_bg,
				walletImg:Images.wallet.ic_block,
				integralData:[],
				fromPage:''
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive',this.fromPage)
				this.$router.back()
			}
		},
		beforeMount() {
			this.integralData = this.$store.state.blockArray
		},
		beforeRouteEnter(to,from,next){
			next(vm=>{
				vm.fromPage = from.name
			})
		}
	}
</script>

<style scoped>
	.bg{
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 7.5rem;
		z-index: -1;
	}
	.van-nav-bar{
		background: #FF7256;
	}
	.van-nav-bar__arrow{
		color: #fff;
		font-size: 1.5rem;
	}
	.van-nav-bar__title{
		color: white;
	}
	.wallet{
		margin:3.5rem 0.625rem 0.625rem;
		text-align: center;
		background: white;
		border-radius: 0.625rem;
		overflow: hidden;
	}
	.wallet01{
		margin-top: 1.25rem;
	}
	.wallet02{
		color: red;
		margin: 0.9375rem auto;
	}
	.list>.van-row{
		background: white;
		margin: 0.625rem 0.625rem;
		border-radius: 0.625rem;
		overflow: hidden;
		text-align: left;
	}
	.list01{
		text-align: center;
	}
	.list01 .van-image{
		margin: 0.3125rem;
	}
	.list02>.van-row:nth-child(1){
		font-size: 0.9375rem;
		height: 1.45625rem;
		line-height: 1.45625rem;
	}
	.list02>.van-row:nth-child(2){
		height: 1.45625rem;
		line-height: 1.45625rem;
	}
	.list02>.van-row:nth-child(2) .van-col:first-child{
		font-size: 0.8125rem;
		opacity: 0.8;
	}
	.list02>.van-row:nth-child(2) .van-col:last-child{
		font-size: 0.875rem;
		color: rgb(255,115,105);
	}
	.list02>.van-row:nth-child(3){
		height: 1.45625rem;
		line-height: 1.45625rem;
		font-size: 0.8125rem;
		opacity: 0.8;
	}
</style>
