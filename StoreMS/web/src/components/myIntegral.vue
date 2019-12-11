<template>
	<div id="myIntegral" :style="bg">
		<div class="content">
			<van-row class="content01">
				<van-col class="col01" span="4">
					<van-image
						round
						width="2.5rem"
						height="2.5rem"
						:src="avatar"
					/>
				</van-col>
				<van-col class="col02" span="13">
					我的积分&nbsp;&nbsp;<span>{{allBlock}}</span>
				</van-col>
				<van-col class="col03" span="7">
					<van-row @click="jumpToWallet()">
						<van-col span="10">
							<van-image :src="jinbi" width="1.5rem" height="1.5rem"/>
						</van-col>
						<van-col span="14">
							积分明细
						</van-col>
					</van-row>
				</van-col>
			</van-row>
			<van-row class="content02">
				<van-col span="24">等值于<span>{{parseFloat(allBlock).toLocaleString()}}</span>元人民币</van-col>
			</van-row>
		</div>
	</div>
</template>

<script>
	import Vue from 'vue'
	import {Row,Col,Image,Button} from 'vant'
	import {Images} from '@/resources/index.js'
	Vue.use(Row).use(Col).use(Image).use(Button)
	export default{
		props:{
			avatar:{
				type:String,
				default:Images.myintergral.ic_default_head
			},
			allBlock:{
				type:Number,
				default:0
			}
		},
		data(){
			return{
				bg:{
					backgroundImage:"url("+Images.myintergral.ic_info_bg+")",
					backgroundRepeat:"no-repeat",
					backgroundSize:"100% 100%",
				},
				jinbi:Images.myintergral.ic_jinbi1,
			}
		},
		methods:{
			jumpToWallet(){
				this.$router.push('/wallet')
			}
		}
	}
</script>

<style scoped>
	#myIntegral{
		width: 100%;
		height: 7rem;
		margin: 0 auto;
	}
	.content{
		width: 95.5%;
		height: 100%;
		margin:0 auto;
		padding-top: 0.48rem;
	}
	.content01{
		height: 3.5rem;
	}
	.content01 .col01 .van-image{
		margin-top: 0.625rem;
	}
	.content01 .col02{
		text-align: left;
		height: 3.5rem;
		line-height: 3.5rem;
		font-size: 0.9375rem;
		color: white;
	}
	.content01 .col02 span{
		font-size: 1.5625rem;
	}
	.content01 .col03>.van-row{
		height: 2rem;
		line-height: 2rem;
		color: white;
		margin-top: 0.75rem;
		background: rgba(255,255,255,0.3);
		border-top-left-radius: 1rem;
		border-bottom-left-radius: 1rem;
	}
	.content01 .col03 .van-image{
		margin-top: 0.25rem;
	}
	.content01 .col03 .van-row .van-col:last-child{
		font-size: 0.75rem;
	}
	.content02{
		height: 2.5rem;
		line-height: 2.5rem;
		color: white;
		font-weight: bold;
		font-size: 0.9375rem;
	}
	.content02 span{
		font-size: 1.5625rem;
	}
</style>
