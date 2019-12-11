<template>
	<div id="userinfo">
		<img :src="bg" class="bg">
		<van-nav-bar
			title="个人资料"
			left-arrow
			fixed
			@click-left="goBack"
			:border="border"
		/>
		
		<van-row class="avatar">
			<van-col span="24">
				<img :src="avatar" @click="showActionSheet">
				<!-- <van-image 
					:src="avatar"
					round
					width="6.25rem"
					height="6.25rem"
					@click="showActionSheet"
				/> -->
			</van-col>
		</van-row>
		
		<van-cell-group class="info">
			<van-cell v-for="(item,index) in getDataList" :key="index"
				:icon="item.image"
				:title="item.title"
				:value="item.content"
				is-link
				@click='to(item.name)'
			>
			</van-cell>
		</van-cell-group>
		
		<van-datetime-picker 
			v-if="dateBar"
			v-model="birthday"
			type="date"
			:min-date="minDate"
			:max-date="maxDate"
			:formatter="formatter"
			@confirm="modifyBirthday"
			@cancel="cancel"
		/>
		
		<van-action-sheet
			v-model="actionShow"
			:actions="actions"
			cancel-text="取消"
			:close-on-click-action="closeonclickaction"
			:safe-area-inset-bottom="safeareainsetbottom"
			@select="onSelect"
			@cancel="onCancel"
		/>
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>

	</div>
</template>

<script>
	import {Images} from '@/resources/index.js'
	import {NavBar,Row,Col,CellGroup,Cell,DatetimePicker,Overlay,Loading,ActionSheet,Toast} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Row.name]:Row,
			[Col.name]:Col,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[DatetimePicker.name]:DatetimePicker,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading,
			[ActionSheet.name]:ActionSheet,
			[Toast.name]:Toast,
		},
		data(){
			return{
				bg:Images.mineView.ic_bg,
				border:false,
				avatar:'',
				birthday:'',
				minDate:new Date(1970,0,1),
				maxDate:new Date(),
				dateBar:false,
				show:false,
				loadingText:'修改中...',
				actionShow:false,
				closeonclickaction:true,
				safeareainsetbottom:true,
				actions:[
					{name:'从相册选取'},
					{name:'拍照'}
				],
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive','mine')
				this.$router.push('/main/mine')
			},
			to(name){
				if (name === 'modifyNickname') {
					this.$router.push('/modifyNickname')
				} else if (name === 'modifyGender') {
					this.$router.push('/modifyGender')
				} else if (name === 'modifyEmail') {
					this.$router.push('/modifyEmail')
				} else if (name === 'modifyBirthday') {
					this.dateBar = true
				} else{
					
				}
			},
			modifyBirthday(value){
				var year = value.getFullYear()
				var month = value.getMonth()+1
				var day = value.getDate()
				var birthday = year+'年'+month+'月'+day+'日'
				this.dateBar = false
				this.show = true
				this.loadingText = '修改中...'
				var url = '/user/birth?user_id='+this.$store.state.userID+'&birth='+birthday+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if (response.data.detail && response.data.detail == 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status == 1) {
						setTimeout(()=>{
							this.loadingText = response.data.detail
							this.$store.commit('birthday',birthday)
							setTimeout(()=>{
								this.show = false
							},500)
						},1000)
					} else{
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
			},
			formatter(type,value){
				if(type === 'year'){
					return `${value}年`
				}else if(type === 'month'){
					return `${value}月`
				}else if(type === 'day'){
					return `${value}日`
				}
				return value
			},
			cancel(){
				this.dateBar = false
			},
			showActionSheet(){
				this.actionShow = true
			},
			onSelect(item){
				if (item.name === '从相册选取') {
					this.actionShow = false
					this.galleryImgs()
				}
				if (item.name === '拍照') {
					this.actionShow = false
					this.getImage()
				}
			},
			onCancel(){
				
			},
			galleryImgs(){
				plus.gallery.pick((path)=>{
					this.imgPreviewnew(path)
				},(e)=>{
					Toast(e.message)
				},{
					filter:"image"
				})
			},
			getImage(){
				var cmr = plus.camera.getCamera()
				cmr.captureImage((p)=>{
					plus.io.resolveLocalFileSystemURL(p,(entry)=>{
						var imgUrl = entry.toLocalURL()
						this.imgPreviewnew(imgUrl)
					})
				},(e)=>{
						Toast(e.message)
				},{
					index:1,
					filename:'_doc/camera/'
				})
			},
			uploadImage(file){
				var url = '/user/upload?token=' + this.$store.state.token
				var formData = new FormData()
				formData.append('file',file)
				formData.append('user_id',this.$store.state.userID)
				this.axios.post(url,formData).then((response)=>{
					if (response.data.detail && response.data.detail == 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if(response.data.status === 1){
						this.avatar = response.data.obj+'?v='+Math.random()
						this.$store.commit('avatar',response.data.obj+'?v='+Math.random())
						Toast('头像上传成功')
					} else{
						Toast('头像上传失败,请重试')
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			imgPreviewnew(file){
				let that = this
				let Orientation
				let img = new Image()
				img.src = file
				img.onload = function(){
					let base64 = that.compress(img,Orientation)
					// console.log(data)
					let blob = that.dataURLtoBlob(base64)
					var file = that.blobToFile(blob,"image.png")
					that.uploadImage(file)
				}
			},
			compress(img,Orientation) {
				let canvas = document.createElement("canvas");
				let ctx = canvas.getContext('2d');
				//瓦片canvas
				let tCanvas = document.createElement("canvas");
				let tctx = tCanvas.getContext("2d");
				let initSize = img.src.length;
				let width = img.width;
				let height = img.height;
				//如果图片大于四百万像素，计算压缩比并将大小压至400万以下
				let ratio;
				if ((ratio = width * height / 4000000) > 1) {
					console.log("大于400万像素")
					ratio = Math.sqrt(ratio);
					width /= ratio;
					height /= ratio;
				} else {
					ratio = 1;
				}
				canvas.width = width;
				canvas.height = height;
				//        铺底色
				ctx.fillStyle = "#fff";
				ctx.fillRect(0, 0, canvas.width, canvas.height);
				//如果图片像素大于100万则使用瓦片绘制
				let count;
				if ((count = width * height / 1000000) > 1) {
					console.log("超过100W像素");
					count = ~~(Math.sqrt(count) + 1); //计算要分成多少块瓦片
					//            计算每块瓦片的宽和高
					let nw = ~~(width / count);
					let nh = ~~(height / count);
					tCanvas.width = nw;
					tCanvas.height = nh;
					for (let i = 0; i < count; i++) {
						for (let j = 0; j < count; j++) {
							tctx.drawImage(img, i * nw * ratio, j * nh * ratio, nw * ratio, nh * ratio, 0, 0, nw, nh);
							ctx.drawImage(tCanvas, i * nw, j * nh, nw, nh);
						}
					}
				} else {
					ctx.drawImage(img, 0, 0, width, height);
				}
				//修复ios上传图片的时候 被旋转的问题
				if(Orientation != "" && Orientation != 1){
					switch(Orientation){
						case 6://需要顺时针（向左）90度旋转
							this.rotateImg(img,'left',canvas);
							break;
						case 8://需要逆时针（向右）90度旋转
							this.rotateImg(img,'right',canvas);
							break;
						case 3://需要180度旋转
							this.rotateImg(img,'right',canvas);//转两次
							this.rotateImg(img,'right',canvas);
							break;
					}
				}
				//进行最小压缩
				let ndata = canvas.toDataURL('image/jpeg', 0.1);
				console.log('压缩前：' + initSize);
				console.log('压缩后：' + ndata.length);
				console.log('压缩率：' + ~~(100 * (initSize - ndata.length) / initSize) + "%");
				tCanvas.width = tCanvas.height = canvas.width = canvas.height = 0;
				return ndata;
			},
			dataURLtoBlob(dataurl){
				var arr = dataurl.split(',')
				var mime = arr[0].match(/:(.*?);/)[1]
				var bstr = atob(arr[1])
				var n = bstr.length
				var u8arr = new Uint8Array(n)
				while (n--){
					u8arr[n] = bstr.charCodeAt(n)
				}
				return new Blob([u8arr],{type:mime})
			},
			blobToFile(blob,fileName){
				blob.lastModifiedDate = new Date()
				blob.name = fileName
				return blob
			}
		},
		mounted() {
			this.avatar = this.$store.state.avatar ? this.$store.state.avatar : Images.mineView.ic_default_head2
			var str = this.$store.state.birthday
			let len1 = str.indexOf('年')
			let len2 = str.indexOf('月')
			let s1 = str.substring(0,len1)
			let s2 = str.substring(len1+1,len2)
			let s3 = str.substring(len2+1,str.length-1)
			this.birthday = new Date(s1,s2-1,s3)
		},
		computed:{
			getDataList(){
				return (
					[
						{image:Images.mineView.ic_info_data[0],title:'昵称',content:this.$store.state.nickname,name:'modifyNickname'},
						{image:Images.mineView.ic_info_data[1],title:'性别',content:this.$store.state.gender,name:'modifyGender'},
						{image:Images.mineView.ic_info_data[2],title:'出生年月',content:this.$store.state.birthday,name:'modifyBirthday'},
						{image:Images.mineView.ic_info_data[3],title:'邮箱',content:this.$store.state.email,name:'modifyEmail'},
						{image:Images.mineView.ic_info_data[4],title:'实名认证',content:'身份证',name:'Autonym'}
					]
				)
			},
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
		background: rgba(0,0,0,0);
	}
	.van-nav-bar__arrow{
		color: #fff;
		font-size: 1.5rem;
	}
	.van-nav-bar__title{
		color: white;
	}
	.avatar{
		margin-top: 3.75rem;
	}
	.avatar img{
		width:6.25rem;
		height:6.25rem;
		border-radius:50%;
		border: 0.0625rem solid #F2F3F5;
	}
	.info{
		margin: 0.625rem 0.625rem 0 0.625rem;
		border-radius: 0.5rem;
		overflow: hidden;
	}
	.info .van-cell{
		text-align: left;
	}
	.van-overlay{
		background: rgba(0,0,0,0);
	}
</style>
