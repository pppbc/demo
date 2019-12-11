<template>
	<div>
		<van-nav-bar
			:title="title"
			right-text="提交"
			left-arrow
			fixed
			@click-left="goBack"
			@click-right="onSubmit"
		/>
		
		<van-card
			:num="data.Quantity"
			:price="data.CurrentUnitPrice"
			:title="data.ProductName"
			:thumb="data.ProductImage"
			currency="Block"
		/>
		
		<van-row class="comment">
			<van-col span="24">
				<van-field
					v-if="title === '退款申请'"
					v-model="reason"
					label="退款原因"
					type="textarea"
					placeholder="请输入退款原因"
					maxlength="500"
					autosize
				/>
				<van-field
				v-else
					v-model="reason"
					label="退货原因"
					type="textarea"
					placeholder="请输入退货原因"
					maxlength="500"
					autosize
				/>
			</van-col>
		</van-row>
		
		<van-row class="price">
			<van-col span="6">退款金额</van-col>
			<van-col span="18">Block {{data.TotalPrice}}</van-col>
		</van-row>
	</div>
</template>

<script>
	import {NavBar,Card,Toast,Field} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Card.name]:Card,
			[Toast.name]:Toast,
			[Field.name]:Field
		},
		data(){
			return {
				title:'',
				data:'',
				reason:'',
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			onSubmit(){
				if (this.reason) {
					var url = ''
					var path = ''
					if (this.title === '退款申请') {
						url = '/order/orderReturn'
						path = '/refundAndReturn/refund/退款'
					}else{
						url = '/order/productReturn'
						path = '/refundAndReturn/return/退货'
					}
					var data = {
						UserID:this.$store.state.userID,
						OrderNo:this.data.OrderNo,
						ProductID:this.data.ProductID,
						Reason:this.reason
					}
					this.axios.post(url,data).then((response)=>{
						if (response.data.status === 1) {
							Toast('已提交申请')
							setTimeout(() => {
								this.$router.push(path)
							}, 1000);
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}else{
					Toast("请填写说明")
				}
			}
		},
		beforeMount() {
			this.title = this.$route.params.title
			this.data = this.$store.state.refundReturnData
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: #000000;
	}
	.van-nav-bar__text{
		color: #000000;
	}
	.van-card{
		margin-top: 2.875rem;
	}
	.van-card__title{
		text-align: left;
	}
	.van-card__price{
		float: left;
	}
	.comment{
		text-align: left;
		margin-top: 0.3125rem;
	}
	.comment .van-field{
		background: #fafafa;
	}
	.price{
		background: #FAFAFA;
		margin-top: 0.3125rem;
		padding: 0.625rem 0;
		font-size: 0.875rem;
	}
	.price .van-col:last-child{
		text-align: left;
		color: red;
	}
</style>
