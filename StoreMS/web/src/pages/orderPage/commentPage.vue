<template>
	<div>
		<van-nav-bar
			title="发表评价"
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
		
		<van-row class="rate">
			<van-col span="4">评分</van-col>
			<van-col span="20"><van-rate v-model="rate"/></van-col>
		</van-row>
		
		<van-row class="comment">
			<van-col span="24">
				<van-field
					v-model="message"
					label="留言"
					type="textarea"
					placeholder="宝贝满足你的期待吗?说说你的使用心得,分享给想买的他们吧"
					label-width="3rem"
					maxlength="500"
					autosize
				/>
			</van-col>
		</van-row>
		
		
	</div>
</template>

<script>
	import {NavBar,Rate,Image,Field,Card,Row,Col,Toast} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Rate.name]:Rate,
			[Image.name]:Image,
			[Field.name]:Field,
			[Card.name]:Card,
			[Row.name]:Row,
			[Col.name]:Col,
			[Toast.name]:Toast
		},
		data(){
			return{
				data:'',
				rate:0,
				message:''
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			onSubmit(){
				var title = ''
				if (this.rate === 1 || this.rate === 2) {
					title = '差评'
				}
				if (this.rate === 3 || this.rate === 4) {
					title = '中评'
				}
				if (this.rate === 5) {
					title = '好评'
				}
				if (this.rate === 0) {
					Toast('请选择你的评分')
					return 
				}
				var url = '/order/addComment?token='+this.$store.state.token
				var data = {
					UserID : this.$store.state.userID,
					OrderNo : this.data.OrderNo,
					ProductID : this.data.ProductID,
					Title : title,
					Content : this.message
				}
				this.axios.post(url,data).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						Toast('评论成功')
						setTimeout(() => {
							this.$router.back()
						}, 1000);
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			}
		},
		beforeMount() {
			this.data = this.$store.state.commentData
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
	.rate{
		background: #fafafa;
		padding: 0.625rem 0;
	}
	.rate .van-col:last-child{
		text-align: left;
	}
	.comment{
		text-align: left;
	}
	.comment .van-field{
		background: #fafafa;
	}
	
</style>
