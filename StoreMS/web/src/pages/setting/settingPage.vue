<template>
	<div id="settingPage">
		<van-nav-bar
			:title="title"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		<div v-if="title === '收货地址'" style="margin-top: 46px;margin-bottom: 80px;overflow: hidden;">
			<van-row class="addressContent" v-for="(item,index) in addressList" :key="index" @click="selectAddr(index)">
				<van-col span="24">
					<van-row class="row1">
						<van-col span="8">{{item.ReceiverName}}</van-col>
						<van-col span="10">{{item.ReceiverMobile}}</van-col>
					</van-row>
					<van-row class="row2">
						<van-col span="24">{{item.ReceiverProvince+' '+item.ReceiverCity+' '+item.ReceiverDistrict+' '+item.DetailAddress}}</van-col>
					</van-row>
					<van-divider />
					<van-row class="row3">
						<van-col v-if="item.DefaultAddress === 1" span="14">
							<van-button class="checkbtn1" tag="span" :icon="checkedIcon">已设为默认</van-button>
						</van-col>
						<van-col v-else span="14">
							<van-button class="checkbtn2" tag="span" :icon="uncheckedIcon" @click="changeDefalutAddr(item.ID,index)">设为默认</van-button>
						</van-col>
						<van-col span="5">
							<van-button class="editbtn" :icon="editIcon" @click="modifyAddr(index)">修改</van-button>
						</van-col>
						<van-col span="5">
							<van-button class="delbtn" :icon="delIcon" @click="delAddr(item.ID,index)">删除</van-button>
						</van-col>
					</van-row>
				</van-col>
			</van-row>
			
			<van-row class="addBtn">
				<van-col span="14" offset="5">
					<van-button class="btn1" type="danger" round @click="addAddr">添加地址</van-button>
				</van-col>
			</van-row>
		</div>
		
		<div v-if="title === '账户与安全'" style="margin-top: 46px;margin-bottom: 80px;overflow: hidden;">
			<van-cell-group class="account">
				<van-cell
					v-for="(item,index) in accountAndSecurityList"
					:key="index"
					:title="item.title"
					:to="item.name"
					is-link
				/>
			</van-cell-group>
		</div>
		<van-overlay :show="show" class-name="overlay" />
		<!-- <van-loading v-if="show" vertical>{{loadingText}}</van-loading> -->
		
	</div>
</template>

<script>
	import {NavBar,Row,Col,CellGroup,Cell,Button,Divider,Overlay,Loading,Toast,Dialog} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Row.name]:Row,
			[Col.name]:Col,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[Button.name]:Button,
			[Divider.name]:Divider,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading,
			[Toast.name]:Toast,
			[Dialog.name]:Dialog
		},
		data(){
			return {
				bg:Images.setting.ic_bg,
				title:'',
				loadingText:'修改中...',
				show:false,
				addressList:[],
				checkedIcon:Images.setting.ic_radio_selected,
				uncheckedIcon:Images.setting.ic_radio_normal,
				editIcon:Images.setting.ic_edit,
				delIcon:Images.setting.ic_delete,
				accountAndSecurityList:[
					{title:'设置登录密码',name:'/loginPassword'},
					{title:'设置支付密码',name:'/tradingPassword'},
				],
				fromPage:''
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			getData(title){
				this.title = title
				if (this.title === '收货地址') {
					this.addressList = this.$store.state.receiverAddressList
				} else if (this.title === '账户与安全'){
					console.log("账号与安全");
				} else {
					console.log("其他界面")
				}
			},
			changeDefalutAddr(ID,index){
				this.show = true
				// this.loadingText = '修改中...'
				var url = '/ship/setDefaultAdd?user_id='+this.$store.state.userID+'&add_id='+ID+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					console.log(response)
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if(response.data.status === 1){
						this.$store.commit("exchangeReceiverItem",index)
						this.$store.commit('currentReceiverAddress')
						setTimeout(()=>{
							Toast(response.data.detail)
							setTimeout(() => {
								this.show = false;
							}, 100)
						},100)
					} else {
						setTimeout(() => {
							Toast(response.data.detail)
							setTimeout(() => {
								this.show = false
							}, 100)
						}, 100)
					}
				}).catch((response)=>{
					setTimeout(()=>{
						Toast('网络请求出错，请稍候再试!')
						this.show = false
					},100)
				})
			},
			delAddr(ID,index){
				Dialog.confirm({
					title:'提示',
					message:'是否确认删除改地址'
				}).then(()=>{
					this.delConfirm(ID,index)
				}).catch(()=>{
					
				})
			},
			delConfirm(ID,index){
				this.show = true
				var url = '/ship/delAdd?user_id='+this.$store.state.userID+'&add_id='+ID+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						this.$store.commit('deleteReceiverItem',index)
						setTimeout(() => {
							Toast('已删除')
							setTimeout(() => {
								this.show = false
							}, 10)
						}, 10)
					} else{
						setTimeout(() => {
							Toast('删除失败，请重试')
							setTimeout(() => {
								this.show = false
							}, 10);
						}, 10);
					}
				}).catch((response)=>{
					setTimeout(() => {
						Toast('网络请求出错，请重试')
						setTimeout(() => {
							this.show = false
						}, 10);
					}, 10);
				})
			},
			modifyAddr(index){
				this.$router.push('/modifyAddress/'+index)
			},
			addAddr(){
				this.$router.push('/addAddress')
			},
			selectAddr(index){
				if (this.fromPage === 'confirmOrder') {
					this.$store.commit('currentReceiverAddress',index)
					this.$router.back()
				}
			}
		},
		beforeRouteEnter(to,from,next){
			next(vm=>{
				vm.getData(to.meta.title)
				vm.fromPage = from.name
			})
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		color: black;
		font-size: 1.2rem;
	}
	.addressContent{
		margin: 0.625rem 0.625rem 0 0.625rem;
		text-align: left;
		background: #FFFFFF;
		border-radius: 0.625rem;
	}
	.row1{
		margin: 0.75rem 0.9375rem 0.625rem ;
		color: rgb(36,36,36);
		font-weight: bold;
	}
	.row1 .van-col:first-child{
		font-size: 1rem;
	}
	.row1 .van-col:last-child{
		font-size: 0.875rem;
	}
	.row2{
		margin: 0 0.9375rem 0.9375rem;
		font-size: 0.875rem;
		color: rgb(36,36,36);
	}
	.row3{
		margin: 1.0625rem 0.9375rem;
	}
	.checkbtn1{
		color: red;
		font-size: 0.75rem;
		margin: 0;
		padding: 0;
		border: none;
		height: 0.75rem;
		line-height: 0.75rem;
	}
	.checkbtn2{
		color: rgb(96,96,96);
		font-size: 0.75rem;
		margin: 0;
		padding: 0;
		border: none;
		height: 0.75rem;
		line-height: 0.75rem;
	}
	.editbtn,.delbtn{
		color: rgb(96,96,96);
		font-size: 0.75rem;
		margin: 0;
		padding: 0;
		border: none;
		height: 0.75rem;
		line-height: 0.75rem;
	}
	.van-button:active::before {
		opacity: 0
	}
	.van-divider{
		margin: 0;
	}
	.addBtn{
		position: fixed;
		left: 0;
		right: 0;
		bottom: 0;
		padding-bottom: 0.9375rem;
		padding-top: 0.625rem;
		width: 100%;
		background: #F5F5F5;
	}
	.btn1{
		width: 100%;
	}
	.account{
		margin: 0.625rem 0.625rem 0 0.625rem;
		border-radius: 0.625rem;
		overflow: hidden;
		text-align: left;
	}
	
	.overlay{
		background-color: rgba(0,0,0,0);
	}
</style>
