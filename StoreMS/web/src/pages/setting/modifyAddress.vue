<template>
	<div >
		<van-nav-bar
			:title="title"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		
		<div v-if="title === '修改地址'" style="margin-top: 46px;">
			<van-address-edit
				:area-list="areaData"
				:address-info="addressInfo"
				:detail-rows="2"
				show-postal
				@save="onSaveEdit"
			/>
		</div>
		
		<div v-if="title === '添加地址'" style="margin-top: 46px;">
			<van-address-edit
				:area-list="areaData"
				:detail-rows="2"
				show-postal
				@save="onSaveAdd"
			/>
		</div>
		
		
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
	</div>
</template>

<script>
	import areaList from '@/utils/area.js'
	import {NavBar,Button,AddressEdit,Toast,Overlay,Loading} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Button.name]:Button,
			[AddressEdit.name]:AddressEdit,
			[Toast.name]:Toast,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading
		},
		data(){
			return {
				show:false,
				loadingText:'',
				title:'',
				index:0,
				areaData:areaList,
				searchResult:[],
				addrID:'',
				receiverName:'',
				receiverMoblie:'',
				receiverProvince:'',
				receiverCity:'',
				receiverCounty:'',
				receiverAddr:'',
				receiverCode:''
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			onSaveAdd(content){
				var data = {
					UserID:this.$store.state.userID,
					ReceiverName:content.name,
					ReceiverMobile:content.tel,
					ReceiverProvince:content.province,
					ReceiverCity:content.city,
					ReceiverDistrict:content.county,
					DetailAddress:content.addressDetail,
					ReceiverZip:content.postalCode,
				}
				this.show = true
				this.loadingText = '正在添加...'
				var url = '/ship/addAddress?token='+this.$store.state.token
				this.axios.post(url,data).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if(response.data.status === 1){
						this.$store.commit("addReceiverItem",response.data.obj)
						setTimeout(() => {
							this.loadingText = '添加成功'
							setTimeout(() => {
								this.show = false
								this.$router.back()
							}, 500)
						}, 1000)
					}else{
						setTimeout(() => {
							Toast('地址添加失败，请重试')
							setTimeout(() => {
								this.show = false
							}, 500)
						}, 1000)
					}
				}).catch((responde)=>{
					setTimeout(() => {
						Toast('网络请求出错，请重试')
						setTimeout(() => {
							this.show = false
						}, 500)
					}, 1000)
				})
			},
			onSaveEdit(content){
				var data = {
					UserID:this.$store.state.userID,
					ID:this.addrID,
					ReceiverName:content.name,
					ReceiverMobile:content.tel,
					ReceiverProvince:content.province,
					ReceiverCity:content.city,
					ReceiverDistrict:content.county,
					DetailAddress:content.addressDetail,
					ReceiverZip:content.postalCode,
				}
				this.show = true
				this.loadingText = '正在修改...'
				var url = '/ship/editAdd?token='+this.$store.state.token
				this.axios.post(url,data).then((response)=>{
					console.log(response)
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if(response.data.status === 1){
						this.$store.commit({type:'modifyReceiverAddressList',data:data,index:this.index})
						setTimeout(() => {
							this.loadingText = '修改成功'
							setTimeout(() => {
								this.show = false
								this.$router.back()
							}, 500)
						}, 1000)
					}else{
						setTimeout(() => {
							Toast('地址修改失败，请重试')
							setTimeout(() => {
								this.show = false
							}, 500)
						}, 1000)
					}
				}).catch((response)=>{
					setTimeout(() => {
						Toast('网络请求出错，请重试')
						setTimeout(() => {
							this.show = false
						}, 500)
					}, 1000)
				})
			}
		},
		beforeRouteEnter(to,from,next){
			next(vm=>{
				vm.title = to.meta.title
				var index = vm.$route.params.index
				vm.index = index
				if (to.meta.title === '修改地址') {
					vm.addrID = vm.$store.state.receiverAddressList[index].ID
					vm.receiverName = vm.$store.state.receiverAddressList[index].ReceiverName
					vm.receiverMoblie = vm.$store.state.receiverAddressList[index].ReceiverMobile
					vm.receiverProvince = vm.$store.state.receiverAddressList[index].ReceiverProvince
					vm.receiverCity = vm.$store.state.receiverAddressList[index].ReceiverCity
					vm.receiverCounty = vm.$store.state.receiverAddressList[index].ReceiverDistrict
					vm.receiverAddr = vm.$store.state.receiverAddressList[index].DetailAddress
					vm.receiverCode = vm.$store.state.receiverAddressList[index].ReceiverZip
				} else{
					vm.addrID = ''
					vm.receiverName = ''
					vm.receiverMoblie = ''
					vm.receiverProvince = ''
					vm.receiverCity = ''
					vm.receiverCounty = ''
					vm.receiverAddr = ''
					vm.receiverCode = ''
				}
			})
		},
		computed:{
			addressInfo(){
				var provinceCode = ''
				var cityCode = ''
				var areaCode = ''
				for (let i in areaList.province_list) {
					if (this.receiverProvince === areaList.province_list[i]) {
						provinceCode = i.toString()
					}
				}
				for (let i in areaList.city_list) {
					var key = i.toString()
					for (var j = 0; j < 2; j++) {
						if (provinceCode[j] === key[j] && this.receiverCity === areaList.city_list[i]) {
							cityCode = key
						}
					}
				}
				for (let i in areaList.county_list) {
					var key = i.toString()
					for (var j = 0; j < 4; j++) {
						if (cityCode[j] === key[j] && this.receiverCounty === areaList.county_list[i]) {
							areaCode = key
						}
					}
				}
				var data = {
					id:this.addrID,
					name:this.receiverName,
					tel:this.receiverMoblie,
					province:this.receiverProvince,
					city:this.receiverCity,
					county:this.receiverCounty,
					addressDetail:this.receiverAddr,
					areaCode:areaCode,
					postalCode:this.receiverCode
				}
				return data
			}
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		color: black;
		font-size: 1.2rem;
	}
	.overlay{
		background-color: rgba(0,0,0,0);
	}
	.van-address-edit{
		text-align: left;
	}
</style>
