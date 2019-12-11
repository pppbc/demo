<template>
	<div id="home">
		<img class="bg" :src="bg" alt="" />
		<van-nav-bar title="布洛克积分商城" :border="border" fixed />
		<div v-if="isonline">
			<my-integral :avatar="avatar" :allBlock="allBlock"></my-integral>

			<div style="overflow-y: hidden;overflow-x: scroll;text-align: left; margin-top: 0.625rem;">
				<ul class="scrollWidth">
					<li v-for="(item, index) in businessArray" :key="index" :style="bgImg(index)">
						<ul class="scrollContent">
							<li>{{ item.PMName }}</li>
							<li>
								<span>{{ item.balance }}</span>
								积分
							</li>
							<li><van-button @click="receivePoints(item)"></van-button></li>
						</ul>
					</li>
				</ul>
			</div>

			<swipe :swipeItems="allActicities"></swipe>

			<van-cell-group><van-cell title="当前可兑换商品" :icon="gift_icon" value="更多" is-link value-class="value_class" /></van-cell-group>

			<van-list v-model="loading" :finished="finished" :finished-text="finishedText" @load="onLoad">
				<double-card :colSpan="colSpan" :doubleCardItems="goods" fromPage="home"></double-card>
			</van-list>
		</div>
		<div v-else class="offline">
			<div>
				<van-image :src="noNetwork" />
				<van-button plain type="info" @click="onRefresh">点击刷新</van-button>
			</div>
		</div>
	</div>
</template>

<script>
import bus from '@/utils/bus.js';
import { NavBar, Toast, Button, Dialog, CellGroup, Cell, List } from 'vant';
import { Images } from '@/resources/index.js';
import myIntegral from '@/components/myIntegral.vue';
import swipe from '@/components/swipe.vue';
import doubleCard from '@/components/double-card.vue';
export default {
	inject:['reload'],
	components: {
		[NavBar.name]: NavBar,
		[Toast.name]: Toast,
		[Button.name]: Button,
		[Dialog.Component.name]: Dialog.Component,
		[CellGroup.name]: CellGroup,
		[Cell.name]: Cell,
		[List.name]: List,
		myIntegral,
		swipe,
		doubleCard
	},
	data() {
		return {
			bg: Images.common.ic_header_bg,
			border: false,
			intrgralData: '',
			businessArray: [],
			goods: [],
			avatar: '',
			allActicities: [],
			gift_icon: Images.home.ic_gift,
			loading: false,
			finished: false,
			finishedText: '没有更多了o(╯□╰)o',
			colSpan: 12,
			isonline:navigator.onLine,
			noNetwork:Images.common.ic_no_network
		};
	},
	methods: {
		//获取用户信息
		getUserInfo() {
			var url = '/user/get?phone=' + this.$store.state.username + '&token=' + this.$store.state.token;
			this.axios
				.get(url)
				.then(response => {
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登录过期，请重新登录');
						this.$router.push('/login');
					} else {
						this.$store.commit('userinfo',response.data.info)
						this.$store.commit('userID', response.data.info.id);
						this.$store.commit('avatar', response.data.info.head ? response.data.head + '?v=' + Math.random() : '');
						this.$store.commit('nickname', response.data.info.name);
						this.$store.commit('gender', response.data.info.sex);
						this.$store.commit('email', response.data.info.email);
						this.$store.commit('birthday', response.data.info.birthday ? response.data.info.birthday : new Date());
						this.avatar = this.$store.state.avatar;
						window.console.log(response)
						window.console.log("获取用户信息？？")
						this.getUserCartDetail();
						this.getUserAddrList();
						this.getCurrentIntegralForUser();
						this.getOtherIntegralForUser();
						this.getBusinessArray();
						this.getAllActivities(this.$store.state.requestActivities);
						this.netApi(this.$store.state.requestHomeGoods);
					}
				})
				.catch(response => {
					Toast('网络请求出错，请稍候再试！');
				});
		},
		// 获取用户的购物车详情
		getUserCartDetail() {
			var url = '/cart/get?user_id=' + this.$store.state.userID + '&token=' + this.$store.state.token;
			this.axios
				.get(url)
				.then(response => {
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录');
						this.$router.push('/login');
					} else if (response.data.status === 1) {
						this.$store.commit('allDatasData', response.data.info);
						window.console.log(response)
						window.console.log("获取用户购物车信息？？")
					}
				})
				.catch(response => {});
		},
		//获取用户收货地址
		getUserAddrList() {
			var url = '/ship/get?user_id=' + this.$store.state.userID + '&token=' + this.$store.state.token;
			this.axios
				.get(url)
				.then(response => {
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录');
						this.$router.push('/login');
					} else {
						this.$store.commit('receiverAddressList', response.data.info);
						this.$store.commit('currentReceiverAddress');
						window.console.log(response)
						window.console.log("获取用户收货地址？？")
					}
				})
				.catch(response => {});
		},
		//获取用户通用积分
		getCurrentIntegralForUser() {
			var url = '/point/getUserpoint?user_id=' + this.$store.state.userID + '&token=' + this.$store.state.token;
			this.axios
				.get(url)
				.then(response => {
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录');
						this.$router.push('/login');
					} else if (response.data.status === 1) {
						this.$store.commit('totalBlock', response.data.info.balance);
						window.console.log("获取用户通用积分？？",balance)
						window.console.log("获取用户通用积分？？")
					}
				})
				.catch(response => {});
		},
		//获取用户其他积分列表
		getOtherIntegralForUser() {
			var url = '/point/getUserpoint?user_id=' + this.$store.state.userID + '&token=' + this.$store.state.token;
			this.axios
				.get(url)
				.then(response => {
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录');
						this.$router.push('/login');
					} else if (response.data.status === 1) {
						this.intrgralData = response.data.info;
						this.$store.commit('saveResult', response.data.info);
						window.console.log(response)
						window.console.log("获取用户积分？？")
					}
				})
				.catch(response => {});
		},
		//获取用户可领取的商家积分列表
		getBusinessArray() {
			var url = '/point/getPointpublish?phone=' + this.$store.state.username + '&token=' + this.$store.state.token;
			this.axios
				.get(url)
				.then(response => {
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录');
						this.$router.push('/login');
					} else if (response.data.status === 1) {
						this.businessArray = response.data.info.slice(0, 8);
						this.$store.commit('businessArray', response.data.info);
						window.console.log(response)
						window.console.log("获取用户可领取积分？？")
					}
				})
				.catch(response => {});
		},
		netApi(count) {
			// var url = '/product/allList?page=' + count;
			var url = '/product/get?token=' + this.$store.state.token
			this.axios
				.get(url)
				.then(response => {
					if (response.data.status === 1 && response.data.info.length > 0) {
						this.$store.commit('requestHomeGoods', 1);
						this.goods = [...this.goods, ...response.data.info];
					}
					window.console.log(response)
					window.console.log("获取netAPI？？")
				})
				.catch(response => {});
		},
		receivePoints(item) {
			Dialog.confirm({
				title: '提示',
				message: '是否确认领取积分'
			})
				.then(() => {
					this.ReceivePoints(item);
				})
				.catch(() => {});
		},
		ReceivePoints(item) {
			var url = '/point/change?id='+item.id+'&user_id='+this.$store.state.userID+'&token='+this.$store.state.token;
			this.axios.get(url)
				// .get(url, {
				// 	params: {
				// 		// username: this.$store.state.username,
				// 		// uid: this.$store.state.userID,
				// 		id: item.ID,
				// 		// point_id: item.PMID,
				// 		// token: this.$store.state.token
				// 	}
				// })
				.then(response => {
					window.console.log(response)
					if (response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录');
						this.$router.push('/login');
					} else if (response.data.status === 1) {
						Toast('领取积分成功');
						this.getBusinessArray();
						this.getOtherIntegralForUser();
					} else {
						Toast('领取积分失败');
					}
				})
				.catch(response => {
					Toast('网络请求出错，请稍后重试！');
				});
		},
		getAllActivities(count) {
			// var url = '/user/getNews?page=' + count;
			// this.axios
			// 	.get(url)
			// 	.then(response => {
			// 		this.$store.commit('requestActivities', 1);
			// 		this.allActicities = response.data.obj.slice(0, 5);
			// 		this.$store.commit('allActicities', response.data.obj);
			// 	})
			// 	.catch(response => {});
		},
		onLoad() {
			setTimeout(() => {
				this.netApi(this.$store.state.requestHomeGoods);
				this.loading = false;
				if (this.$store.state.requestHomeGoods >= 5) {
					this.finished = true;
				}
			}, 1000);
		},	
		updateOnlineStatus(e){
			const {type} = e
			this.isonline = type === 'online'
		},
		onRefresh(){
			this.reload()
		}
	},
	beforeMount() {
		this.$store.commit('requestHomeGoods', 0);
		this.$store.commit('requestActivities', 0);
		this.$store.commit('username', localStorage.getItem('username'));
		this.$store.commit('token', localStorage.getItem('token'));
		this.getUserInfo();
	},
	mounted() {
		window.addEventListener('online',this.updateOnlineStatus)
		window.addEventListener('offline',this.updateOnlineStatus)
		bus.$on('receivePoints', item => {
			this.receivePoints(item);
		}),
			bus.$on('getAllActivities', msg => {
				this.getAllActivities(this.$store.state.requestActivities);
			});
	},
	beforeDestroy() {
		window.removeEventListener('online',this.updateOnlineStatus)
		window.removeEventListener('offline',this.updateOnlineStatus)
	},	
	computed: {
		allBlock() {
			return this.$store.getters.allBlock;
		},
		bgImg() {
			return function(index) {
				return {
					backgroundImage: 'url(' + Images.home.ic_integral_data[index] + ')',
					backgroundRepeat: 'no-repeat',
					backgroundSize: '100% 100%'
				};
			};
		}
	},
	watch:{
		isonline(){
			this.isonline = navigator.onLine
		}
	}
};
</script>

<style scoped>
::-webkit-scrollbar {
	display: none;
}
.bg {
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	height: 6.5rem;
	z-index: -1;
}
.van-nav-bar {
	background: url(../assets/Images/Home/header_bg.png);
}
.van-nav-bar__title {
	color: #fff;
	font-size: 1.125rem;
}
.content {
	width: 100%;
	height: 62.5rem;
	background: rgba(0, 0, 0, 0);
}
.scrollWidth {
	/* width: 51.125rem; */
	width: auto;
	white-space: nowrap;
	margin-left: 0.625rem;
}
.scrollWidth > li {
	list-style: none;
	display: inline-block;
	/* float: left; */
	margin-right: 0.3125rem;
	width: 5.78125rem;
	height: 5.78125rem;
	border-radius: 0.3125rem;
}

.scrollContent li {
	list-style: none;
	margin: 0.3rem;
	color: white;
}
.scrollContent .van-button {
	width: 3.125rem;
	height: 0.9375rem;
	border: none;
	background: url('../assets/Images/Home/input.png') no-repeat center;
	background-size: 100% 100%;
}
.van-cell-group {
	margin: 0.625rem 0.625rem 0 0.625rem;
	background: rgba(0, 0, 0, 0);
}
.van-cell {
	background: #ffa500;
	border-radius: 0.3125rem;
	height: 2.5rem;
	color: white;
	text-align: left;
}
.van-cell__right-icon {
	color: white;
}
.value_class {
	color: white;
}
.offline{
	width: 100%;
	margin: 0 auto;
}
.offline div{
	width: 50%;
	position: absolute;
	left: 50%;
	top: 20%;
}
.offline .van-image{
	width: 100%;
	height: auto;
	position: relative;
	left: -50%;
}
.offline .van-button{
	width: 6.25rem;
	position: relative;
	left: -50%;
	margin-top: 0.625rem;
}
</style>
