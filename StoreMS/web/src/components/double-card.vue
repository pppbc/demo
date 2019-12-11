<template>
	<van-row id="double-card">
		<van-col :span="colSpan" v-for="(card,index) in doubleCardItems" :key="index" @click="getDetail(card)">
			<van-row>
				<van-col class="content" span="22">
					<van-image width="100%" height="7.5rem" :src="card.MainImage"/>
					<van-row class="prdName">
						<van-col span="24">{{card.Name}}</van-col>
					</van-row>
					<van-row class="prdPrice">
						<van-col span="4">
							<van-image width="0.9375rem" height="0.75rem" :src="jinbi"/>
						</van-col>
						<van-col class="price" span="20"><span>{{card.Price}}</span>积分</van-col>
					</van-row>
				</van-col>
			</van-row>
		</van-col>
	</van-row>
</template>

<script>
	import {Card,Row,Col,Image,Button} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[Card.name]:Card,
			[Row.name]:Row,
			[Col.name]:Col,
			[Image.name]:Image,
			[Button.name]:Button
		},
		props:{
			doubleCardItems:Array,
			colSpan:Number,
			fromPage:String
		},
		data(){
			return{
				centered:true,
				jinbi:Images.home.ic_jinbi2,
			}
		},
		methods:{
			getDetail(card){
				this.$store.commit("fromPage",this.fromPage)
				this.$store.commit("goodsDetail",card)
				this.$router.push('/goodsDetail')
			}
		}
	}
</script>

<style scoped>
	#double-card{
		margin: 0.625rem 0.3125rem 0;
	}
	.content{
		margin: 0 0.45rem 0.625rem 0.45rem;
		border-radius: 0.5rem;
		overflow: hidden;
		background: white;
	}
	.prdName{
		margin: 0 auto;
		padding: 0;
		text-align: left;
		overflow: hidden;
		height: 2.32rem;
		font-size: 0.75rem;
		color: #666666;
	}
	.prdName .van-col{
		padding:0 0.3125rem;
	}
	.prdPrice{
		height: 1.875rem;
	}
	.price{
		text-align: left;
		color: #999999;
		font-size: 0.6875rem;
	}
	.price span{
		color: red;
		font-size: 1.125rem;
		margin-right: 0.3125rem;
	}
</style>
