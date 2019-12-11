<template>
	<van-row class="main">
		<van-col span="24">
			<van-row class="firstLine">
				<van-col span="10">{{title}}</van-col>
				<van-col span="14">{{time}}</van-col>
			</van-row>
			<van-divider />
			<van-row class="secondLine" @click="$emit('orderDetail','订单详情')">
				<van-col span="7">
					<van-image :src="image" width="100%" height="5rem"/>
				</van-col>
				<van-col span="10">{{name}}</van-col>
				<van-col span="7">
					<div>共{{goodsNum}}件</div>
					<div>{{totalPrice}}积分</div>
				</van-col>
			</van-row>
			<van-divider />
			<van-row class="thirdLine">
				<van-col span="24">
					<van-row v-if="title === '待付款'">
						<van-col span="6" offset="12">
							<van-button type="warning" round @click="$emit('cancelOrder','取消订单')">取消订单</van-button>
						</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('orderDetail','去付款')">去付款</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '待发货'">
						<van-col class="orderNo" span="12">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="warning" round @click="$emit('refund','退款')">退款</van-button>
						</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('confirmReceipt','确认收货')">确认收货</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '待收货'">
						<van-col class="orderNo" span="12">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="warning" round @click="$emit('returnOrder','退货')">退货</van-button>
						</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('confirmReceipt','确认收货')">确认收货</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '待评价'">
						<van-col class="orderNo" span="12">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="warning" round @click="$emit('returnOrder','退货')">退货</van-button>
						</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('comment','去评价')">去评价</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '退款中'">
						<van-col class="orderNo" span="18">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('cancelRefund','取消退款')">取消退款</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '退款成功' || title === '退款失败'">
						<van-col class="orderNo" span="18">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('orderDetail','查看详情')">查看详情</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '退货申请中'">
						<van-col class="orderNo" span="18">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('orderDetail','查看详情')">查看详情</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '退货申请成功' || title === '退货申请失败'">
						<van-col class="orderNo" span="18">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('orderDetail','查看详情')">查看详情</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '等待卖家收货'">
						<van-col class="orderNo" span="18">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('orderDetail','查看详情')">查看详情</van-button>
						</van-col>
					</van-row>
					<van-row v-else-if="title === '卖家同意退款' || title === '卖家拒绝退款'">
						<van-col class="orderNo" span="18">订单号:{{orderNo}}</van-col>
						<van-col span="6">
							<van-button type="danger" round @click="$emit('orderDetail','查看详情')">查看详情</van-button>
						</van-col>
					</van-row>
					<van-row v-else>
						<van-col span="6" offset="18">
							<van-button type="danger" round @click="$emit('delete','删除订单')">删除订单</van-button>
						</van-col>
					</van-row>
				</van-col>
			
			</van-row>
		</van-col>
	</van-row>
</template>

<script>
	import {Row,Col,Image,Button,Divider} from 'vant'
	export default{
		components:{
			[Row.name]:Row,
			[Col.name]:Col,
			[Image.name]:Image,
			[Button.name]:Button,
			[Divider.name]:Divider
		},
		props:{
			title:String,
			time:String,
			image:String,
			name:String,
			goodsNum:Number,
			totalPrice:Number,
			orderNo:Number
		},
		methods:{
			
		}
	}
</script>

<style scoped>
	.main{
		width: 95%;
		margin: 10px auto;
		background: white;
		border-radius: 0.625rem;
		overflow: hidden;
	}
	.van-divider{
		margin: 0 0.625rem;
		border:0.1px solid #F0F0F0;
	}
	.firstLine{
		margin: 0.5rem 0.625rem;
	}
	.firstLine .van-col:first-child{
		text-align: left;
		font-size: 0.9375rem;
		font-weight: bold;
		opacity: 0.8;
	}
	.firstLine .van-col:last-child{
		text-align: right;
		font-size: 0.8125rem;
		object-fit: 0.7;
	}
	.secondLine{
		margin: 0.5rem 0.625rem;
		font-size: 0.875rem;
		height: 5rem;
		text-align: left;
		overflow: hidden;
	}
	.secondLine .van-col:first-child{
		padding-right: 0.3125rem;
	}
	.secondLine .van-col:last-child{
		padding-left: 0.3125rem;
		text-align: center;
		height: 5rem;
		line-height: 2.5rem;
		font-size: 0.75rem;
	}
	.secondLine .van-col:last-child div{
		height: 50%;
	}
	.thirdLine{
		margin: 0.5rem 0.625rem;
	}
	.thirdLine .orderNo{
		text-align: left;
		font-size: 0.8125rem;
		opacity: 0.7;
	}
	.thirdLine .van-button{
		margin: 0;
		height: 1.5rem;
		line-height: 1.5rem;
		font-size: 0.6875rem;
	}
</style>
