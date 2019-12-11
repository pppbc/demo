<template>
	<div>
		<van-nav-bar
			title="确认付款"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		
		<van-row class="container">
			<van-col span="24">
				<van-row class="totalPrice">
					<van-col span="24"><span>Block </span>{{totalPrice}}</van-col>
				</van-row>
				<van-row class="text">
					<van-col span="24">请选择您的支付方式</van-col>
				</van-row>
				<van-row class="list">
					<van-col span="24">
						<van-collapse v-model="activeName" accordion>
							<van-collapse-item name="1" class="collapseItem1">
								<div slot="title">
									<van-checkbox v-model="checked1">使用通用积分支付</van-checkbox>
								</div>
								<van-row>
									<van-col span="24">当前可用通用积分 {{$store.state.totalBlock}} Block</van-col>
								</van-row>
							</van-collapse-item>
							<van-collapse-item name="2" class="collapseItem2">
								<div slot="title">
									<van-checkbox v-model="checked2">使用自定义积分支付</van-checkbox>
								</div>
								<van-row v-for="(item,index) in blockArr" :key="index" class="blocklist">
									<van-col span="20">{{item.PointName}} 积分余额：{{item.Balance}} 约合：{{item.Balance*item.PointFate}}Block</van-col>
									<van-col span="2">
										<van-icon v-show="index !== 0" name="upgrade" @click="upBtn(index)" />
									</van-col>
									<van-col span="2">
										<van-icon v-show="index !== blockArr.length-1" :name="downgrade" @click="downBtn(index)"/>
									</van-col>
								</van-row>
							</van-collapse-item>
							<van-collapse-item name="3" class="collapseItem3">
								<div slot="title">
									<van-checkbox v-model="checked3">按照积分过期时间自动扣款</van-checkbox>
								</div>
								<van-row>
									<van-col span="24">按照积分过期时间自动扣款</van-col>
								</van-row>
							</van-collapse-item>
						</van-collapse>
					</van-col>
				</van-row>
			</van-col>
		</van-row>
		<van-button type="danger" round @click="onSubmit">确认支付</van-button>
		<van-action-sheet v-model="showActionSheet" title="输入支付密码" @closed="closed">
			<van-password-input
				:value="password"
				:focused="showKeyboard"
				@focus="showKeyboard = true"
			/>
			<van-number-keyboard
				:show="showKeyboard"
				@input="onInput"
				@delete="onDelete"
				:safe-area-inset-bottom="true"
			/>
		</van-action-sheet>
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
	</div>
</template>

<script>
	import {NavBar,Row,Col,Checkbox,Collapse,CollapseItem,Button,Icon,ActionSheet,PasswordInput,NumberKeyboard,Toast,Overlay,Loading} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Row.name]:Row,
			[Col.name]:Col,
			[Checkbox.name]:Checkbox,
			[Collapse.name]:Collapse,
			[CollapseItem.name]:CollapseItem,
			[Button.name]:Button,
			[Icon.name]:Icon,
			[ActionSheet.name]:ActionSheet,
			[PasswordInput.name]:PasswordInput,
			[NumberKeyboard.name]:NumberKeyboard,
			[Toast.name]:Toast,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading
		},
		data(){
			return{
				data:'',
				blockArr:[],
				activeName:'1',
				checked1:true,
				checked2:false,
				checked3:false,
				downgrade:Images.shop.ic_down,
				showActionSheet:false,
				password:'',
				showKeyboard:true,
				show:false,
				loadingText:'正在支付...'
			}
		},
		methods:{
			goBack(){
				this.$router.push('/orderPage/unpaid/待付款')
			},
			upBtn(index){
				var arr = this.blockArr
				var item = arr[index]
				arr[index] = arr[index-1]
				arr[index-1] = item
				this.blockArr.splice(index,1,arr[index])
				this.blockArr.splice(index-1,1,arr[index-1])
			},
			downBtn(index){
				var arr = this.blockArr
				var item = arr[index]
				arr[index] = arr[index+1]
				arr[index+1] = item
				this.blockArr.splice(index,1,arr[index])
				this.blockArr.splice(index+1,1,arr[index+1])
			},
			onSubmit(){
				this.judgePaymentPW()
			},
			onInput(key){
				this.password = (this.password + key).slice(0,6)
				if(this.password.length === 6){
					this.$nextTick(()=>{
						this.showActionSheet = false
						this.show = true
						this.loadingText = '正在支付...'
						var data = {
							UserID:this.$store.state.userID,
							OrderNo:this.data.OrderNo,
							TotalPay:this.totalPrice,
							Password:this.password
						}
						if (this.activeName === '1') {
							data.Type = 1
						}
						if (this.activeName === '2') {
							data.Type = 3
							var arr = []
							for (let item of this.blockArr) {
								arr.push(item.ID)
							}
							data.list = arr
						}
						if (this.activeName === '3') {
							data.Type = 2
						}
						var url = '/point/payWithPoint?token='+this.$store.state.token
						this.axios.post(url,data).then((response)=>{
							if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
								Toast('登陆过期，请重新登录')
								this.$router.push('/login')
							}else if (response.data.status === 1) {
								setTimeout(()=>{
									Toast('支付成功')
									setTimeout(()=>{
										this.show = false
										this.$router.push('/orderPage/unshipped/待发货')
									},500)
								},1000)
							}else{
								setTimeout(()=>{
									Toast(response.data.detail)
									this.show = false
								},1000)
							}
						}).catch((response)=>{
							setTimeout(()=>{
								Toast('网络请求出错，请稍候再试!')
								this.show = false
							},1000)
						})
					})
				}
			},
			onDelete(){
				this.password = this.password.slice(0,this.password.length - 1)
			},
			closed(){
				this.password = ''
			},
			judgePaymentPW(){
				var url = '/user/getpaypwd?user_id='+this.$store.state.userID+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						this.showActionSheet = true
					} else{
						this.$router.push('/tradingPassword')
					}
				}).catch((response)=>{
					Toast('网络请求出错，请稍候再试!')
				})
			}
		},
		computed:{
			totalPrice(){
				if (this.data.length) {
					var totalPrice = 0
					for (let item of this.data) {
						totalPrice += item.TotalPrice
					}
					return totalPrice 
				} else{
					return this.data.TotalPrice
				}
			}
		},
		watch:{
			activeName(){
				if (this.activeName === '1') {
					this.checked1 = true
					this.checked2 = this.checked3 = false
				}
				if (this.activeName === '2') {
					this.checked2 = true
					this.checked1 = this.checked3 = false
				}
				if (this.activeName === '3') {
					this.checked3 = true
					this.checked1 = this.checked2 = false
				}
			}
		},
		beforeMount() {
			this.data = this.$store.state.paymentData
			this.blockArr = this.$store.state.blockArray.filter(val=>{
				return val.Balance !== 0
			})
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: #000000;
	}
	.container{
		margin-top: 2.875rem;
		background: white;
	}
	.totalPrice{
		margin-top: 0.625rem;
		font-size: 1.125rem;
	}
	.totalPrice span{
		font-size: 0.875rem;
		color: #999999;
	}
	.text{
		margin-top: 0.625rem;
		font-size: 0.6875rem;
		opacity: 0.8;
	}
	.list{
		margin-top: 0.625rem;
	}
	.btn{
		margin-top: 0.625rem;
		
	}
	.van-button{
		margin: 2rem auto;
		width: 50%;
	}
	.blocklist{
		text-align: left;
		height: 1.5rem;
		line-height: 1.5rem;
	}
	.blocklist .van-icon{
		font-size: 1.5rem;
	}
	.van-number-keyboard{
		position: relative;
		margin-top: 0.625rem;
	}
	.overlay{
		background: rgba(0,0,0,0);
	}
	.van-loading{
		position: absolute;
		left: 40%;
		top: 45%;
	}
</style>
