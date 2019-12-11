<template>
	<div>
		<van-nav-bar :title="title" left-arrow fixed @click-left="goBack" />
		<div style="overflow-x: hidden;overflow-y: scroll;margin-top: 2.875rem;">
			<div id="convo">
				<ul class="chat-thread">
					<li class="chat" v-for="(item, index) in msglist" :key="index" :class="item.ismine ? 'mine' : 'other'">
						<div><img class="" :src="item.user.avatar || '@/assets/Images/Mine/head1.png'" /></div>
						<span></span>
						<div class="content">
							<div v-if="item.msg.media === 1" v-text="item.msg.content"></div>
							<img class="pic" v-if="item.msg.media === 4" :src="item.msg.url" />
							<div v-if="item.msg.media === 3" @click="playAudio(item.msg.url)">
								<img class="audio" :src="audiom" />
								<span v-text="item.msg.amount"></span>
							</div>
						</div>
					</li>
				</ul>
			</div>
		</div>
		<van-row class="flex-container">
			<van-col class="item-1" span="3" @click="txtstat = 'audio'" v-if="txtstat === 'kbord'">
				<img :src="audio" />
			</van-col>
			<van-col class="item-1" span="3" @click="txtstat = 'kbord'" v-if="txtstat === 'audio'">
				<img :src="kbord" />
			</van-col>
			<van-col class="item-2" span="15" v-if="txtstat === 'kbord'">
				<van-field v-model="txtmsg" placeholder="在这里写点啥" border/>
			</van-col>
			<van-col class="item-2" span="15" v-if="txtstat === 'audio'">
				<input type="button" value="按住 说话" @touchstart="startrecorder" @touchend="stoprecorder"/>
			</van-col>
			<van-col class="item-3" span="3" @click="panelstat = 'doutu'">
				<img :src="smile" />
			</van-col>
			<van-col class="item-4" span="3" v-if="!txtmsg" @click="panelstat = 'more'">
				<img :src="more" />
			</van-col>
			<van-col class="item-4" span="3" v-if="!!txtmsg" @click="sendtxtmsg(txtmsg)">
				<img :src="send" />
			</van-col>
		</van-row>
		<div id="panels" style="display: flex;">
			<div v-if="panelstat === 'doutu'">
				<div class="doutures">
					<div class="res" v-for="(item, index) in doutu.choosed.assets" @click="sendpicmsg(item)" :key="index"><img :class="doutu.choosed.size || 'small'" :src="item" /></div>
				</div>
				<div class="doutupkg">
					<div class="pkg" v-for="(item, index) in doutu.packages" @click="doutu.choosed = item" :key="index"><img :class="item.size || 'small'" :src="item.icon" /></div>
				</div>
			</div>
			<div v-if="panelstat === 'more'" class="plugins">
				<div class="plugin" @click="dispatchplugin(item)" v-for="(item, index) in plugins" :key="index">
					<label :for="item.id"><img :src="item.icon" /></label>
					<div v-html="item.slot"></div>
					<p v-text="item.name"></p>
				</div>
			</div>
		</div>
		<audio id="audio" style="display: none;"></audio>
		<audio id="audio4play" style="display: none;"></audio>
		<audio id="video" style="display: none;"></audio>
		<audio id="video4play" style="display: none;"></audio>
		<div id="sound-alert" class="rprogress" v-show="showprocess">
			<div class="rschedule"></div>
			<div class="r-sigh">!</div>
			<div id="audio-tips" class="rsalert">手指上滑，取消发送</div>
		</div>
	</div>
</template>

<script>
import { NavBar, Toast, Row, Col, Field} from 'vant';
import { Images } from '@/resources/index.js';
import emoj from '@/assets/plugins/doutu/emoj/info.json';
import mkgif from '@/assets/plugins/doutu/mkgif/info.json';
export default {
	components: {
		[NavBar.name]: NavBar,
		[Toast.name]: Toast,
		[Row.name]: Row,
		[Col.name]: Col,
		[Field.name]: Field
	},
	data() {
		return {
			data: this.$store.state.chatting,
			audiom: Images.chat.ic_audiom,
			audio: Images.chat.ic_audio,
			kbord: Images.chat.ic_kbord,
			smile: Images.chat.ic_smile,
			more: Images.chat.ic_more,
			send: Images.chat.ic_send,
			usermap: {},
			webSocket: {},
			txtmsg: '',
			panelstat: 'kbord',
			txtstat: 'kbord',
			showprocess: false,
			recorder: {},
			duration: 0,
			timer: 0,
			msglist: [],
			msgcontext: {
				dstid: 10,
				cmd: 10,
				userid: this.$store.state.userID
			},
			pkg: [mkgif, emoj],
			doutu: {
				config: {
					baseurl: '@/assets/plugins/doutu/',
					pkgids: ['mkgif', 'emoj']
				},
				packages: [],
				choosed: { pkgid: 'emoj', assets: [], size: 'small' }
			},
			plugins: [
				{
					icon: Images.chat.ic_upload,
					name: '照片',
					id: 'upload',
					slot: "<input id='upload' accept='image/gif,image/jepg,image/png' type='file' onchange='upload(this)' class='upload' hidden/>"
				},
				{
					icon: Images.chat.ic_camera,
					name: '拍照',
					id: 'camera',
					slot: "<input id='camera' accept='image/*' capture='camera' type='file' onchange='upload(this)' class='upload' hidden/>"
				},
				{
					icon: Images.chat.ic_audiocall,
					name: '语音',
					id: 'audiocall'
				},
				{
					icon: Images.chat.ic_videocall,
					name: '视频',
					id: 'videocall'
				},
				{
					icon: Images.chat.ic_redpackage,
					name: '红包',
					id: 'redpackage'
				},
				{
					icon: Images.chat.ic_exchange,
					name: '转账',
					id: 'exchange'
				},
				{
					icon: Images.chat.ic_address,
					name: '地址',
					id: 'address'
				},
				{
					icon: Images.chat.ic_person,
					name: '名片',
					id: 'person'
				}
			]
		};
	},
	methods: {
		goBack() {
			this.$router.back();
		},
		playAudio(url) {
			document.getElementById('audio4play').src = url;
			document.getElementById('audio4play').play();
		},
		startrecorder() {
			var audioTarget = document.getElementById('audio');
			var types = ['video/webm', 'audio/webm', 'video/webm;codecs-vp8', 'video/webm;codecs=daala', 'video/webm;codecs=h264', 'audio/webm;codecs=opus', 'video/mpeg'];
			var supporttype = '';
			for (var i in types) {
				if (MediaRecorder.isTypeSupported(types[i])) {
					supporttype = types[i];
				}
			}
			if (!supporttype) {
				Toast('编码不支持');
				return;
			}
			this.duration = new Date().getTime();
			navigator.mediaDevices
				.getUserMedia({ audio: true, video: false })
				.then(stream => {
					this.showprocess = true;
					this.recorder = new MediaRecorder(stream);
					audioTarget.srcObject = stream;
					this.recorder.ondataavailable = event => {
						this.uploadblob('attach/upload', event.data, '.mp3');
						this.$nextTick(() => {
							stream.getTracks().forEach(track => {
								track.stop();
							});
							this.showprocess = false;
						});
					};
					this.recorder.start();
				})
				.catch(err => {
					Toast(err);
					this.showprocess = false;
				});
		},
		stoprecorder() {
			if (typeof this.recorder.stop === 'function') {
				this.recorder.stop();
			}
			this.showprocess = false;
		},
		sendtxtmsg(txt) {
			var msg = this.createmsgcontext();
			msg.media = 1;
			msg.content = txt;
			this.showmsg(this.$store.state.userinfo, msg);
			// this.webSocket.send(JSON.stringify(msg))
		},
		sendpicmsg(picurl) {
			var msg = this.createmsgcontext();
			msg.media = 4;
			msg.url = picurl;
			this.showmsg(this.$store.state.userinfo, msg);
			// this.webSocket.send(JSON.stringify(msg))
		},
		createmsgcontext() {
			return JSON.parse(JSON.stringify(this.msgcontext));
		},
		dispatchplugin(item) {
			switch (item.id) {
				case 'upload':
					break;
				case 'camera':
					break;
				default:
					Toast('该功能暂不支持');
					break;
			}
		},
		upload(dom) {
			this.uploadfile('attach/upload', dom);
		},
		uploadfile(uri, dom) {
			var formdata = new FormData();
			formdata.append('userid', this.$store.state.userID);
			formdata.append('file', dom.files[0]);
			this.axios
				.post(uri, formdata)
				.then(response => {})
				.catch(err => {});
		},
		loaddoutures() {
			var res = [];
			var config = this.doutu.config;
			for (var i in config.pkgids) {
				res[config.pkgids[i]] = require(`@/assets/plugins/doutu/${config.pkgids[i]}/info.json`);
			}
			for (var id in res) {
				for (var j in res[id].assets) {
					res[id].assets[j] = require(`@/assets/plugins/doutu/${id}/${j}.gif`);
				}
				res[id].icon = require(`@/assets/plugins/doutu/${id}/icon.gif`);
				this.doutu.packages.push(res[id]);
				if (this.doutu.choosed.pkgid === res[id].id) {
					this.doutu.choosed.assets = res[id].assets;
				}
			}
		},
		initwebsocket() {
			var url = 'ws://' + location.host + '/chat?id=' + this.$store.state.userID + '&token=' + this.$store.state.token;
			this.webSocket = new WebSocket(url);
			this.webSocket.onmessage = evt => {
				if (evt.data.indexOf('}') > -1) {
					this.onmessage(JSON.parse(evt.data));
				} else {
					console.log('recv<==' + evt.dat);
				}
			};
			this.webSocket.onclose = evt => {
				console.log(evt.data);
			};
			this.webSocket.onerror = evt => {
				console.log(evt.data);
			};
		},
		onmessage(data) {
			this.loaduserinfo(data.userid, user => {
				this.showmsg(user, data);
			});
		},
		loaduserinfo(userid, cb) {
			userid = '' + userid;
			var userinfo = this.usermap[userid];
			if (!userinfo) {
				this.axios.post('user/find', { id: userid }).then(response => {
					cb(response.data);
					this.usermap[userid] = response.data;
				});
			} else {
				cb(userinfo);
			}
		},
		showmsg(user, msg) {
			var data = {};
			data.ismine = this.$store.state.userID === msg.userid;
			data.user = user;
			data.msg = msg;
			this.msglist = this.msglist.concat(data);
			this.reset();
			this.timer = setTimeout(() => {
				window.scrollTo(0, document.getElementById('convo').offsetHeight);
				clearTimeout(this.timer);
			}, 100);
		},
		reset() {
			this.panelstat = 'kbord';
			this.txtstat = 'kbord';
			this.txtmsg = '';
		},
		uploadblob(uri, blob, filetype) {
			var formdata = new FormData();
			formdata.append('filetype', filetype);
			formdata.append('userid', this.$store.state.userID);
			formdata.append('file', blob);
			this.axios
				.post(uri, formdata)
				.then(response => {
					if (response.data.status === 1) {
						var duration = Math.ceil((new Date().getTime() - this.duration) / 1000);
						this.sendaudiomsg(response.data, duration);
					}
				})
				.catch(response => {});
		},
		sendaudiomsg(url, num) {
			var msg = this.createmsgcontext();
			msg.media = 3;
			msg.url = url;
			msg.amount = num;
			this.showmsg(this.$store.state.userinfo, msg);
			// this.webSocket.send(JSON.stringify(msg))
		}
	},
	computed: {
		title() {
			var title = this.data.title;
			if (title.length > 6) {
				title = title.substr(0, 6) + '...';
			}
			return title;
		}
	},
	beforeMount() {
		this.loaddoutures();
		// this.initwebsocket()
	}
};
</script>

<style scoped>
.van-nav-bar__arrow {
	font-size: 1.5rem;
	color: #000;
}
.van-nav-bar__title {
	font-size: 1.125rem;
}
#convo {
	/* margin-top: 2.875rem; */
}
.flex-container {
	/* display: flex; */
	/* flex-direction: row; */
	width: 100%;
	padding-top: 0.5rem;
	position: fixed;
	left: 0;
	bottom: 0px;
	background: #F7F8FA;
}
.item-2 .van-field{
	padding: 0.25rem;
}
.item-2 input{
	background: #FFFFFF;
	text-align: center;
	width: 100%;
	padding: 0.25rem 0;
	border: 0 none;
}
/* .item-1 {
	height: 3.125rem;
	width: 3.125rem;
	padding: 0.3125rem;
}
.item-2 {
	height: 2.5rem;
	border: 1px solid #f0f0f0;
	width: 100%;
}
.item-2 input {
	width: 100%;
	height: 100%;
}
.txt {
	margin-right: auto;
}
.item-3 {
	height: 3.125rem;
	width: 3.125rem;
	padding: 0.3125rem;
}
.item-4 {
	height: 3.125rem;
	width: 3.125rem;
	padding: 0.3125rem;
} */
li.chat {
	justify-content: flex-start;
	align-items: flex-start;
	display: flex;
}
.chat.other {
	flex-direction: row;
}
.chat.mine {
	flex-direction: row-reverse;
}
img.avatar {
	width: 4rem;
	height: 4rem;
}
.other .avatar {
	margin-left: 0.625rem;
}
.mine .avatar {
	margin-right: 0.625rem;
}
.other span {
	border: 10px solid;
	border-color: transparent #ffffff transparent transparent;
	margin-top: 0.625rem;
}
.mine span {
	border: 10px solid;
	border-color: transparent transparent transparent #32cd32;
	margin-top: 0.625rem;
}
.other > .content {
	background-color: #ffffff;
}
.mine > .content {
	background: #32cd32;
}
div.content {
	min-width: 3.75rem;
	clear: both;
	display: inline-block;
	padding: 1rem 1rem 1rem 0.625rem;
	margin: 0 0 20px 0;
	font: 16px/20px 'Noto Sans', sans-serif;
	border-radius: 0.625rem;
	min-height: 4rem;
}
.content > img.pic {
	width: 100%;
	color: white;
}
#panels {
	background: #ffffff;
	display: flex;
	position: fixed;
	bottom: 3.125rem;
}
.doutures {
	flex-direction: row;
	flex-wrap: wrap;
	display: flex;
}
.doutures img {
	margin: 0.625rem;
}
.doutupkg {
	flex-direction: row;
	flex-wrap: wrap;
	display: flex;
	border-top: 0.03125rem solid #f0f0f0;
}
.plugins {
	flex-direction: row;
	flex-wrap: wrap;
	display: flex;
}
.plugin {
	padding: 0.625rem 0.625rem 0.625rem 1.5rem;
	margin: 0 0.625rem;
}
.plugin img {
	width: 2rem;
}
.plugin p {
	text-align: center;
	font-size: 0.875rem;
}
.doutupkg img {
	width: 2rem;
	height: 2rem;
	margin: 0.3125rem;
}
.upload {
	width: 4rem;
	height: 4rem;
	position: absolute;
	top: 0.0625rem;
	opacity: 0;
}
.tagicon {
	width: 2rem;
	height: 2rem;
}
.small {
	width: 2rem;
	height: 2rem;
}
.middle {
	width: 4rem;
	height: 4rem;
}
.large {
	width: 6rem;
	height: 6rem;
}
.res img {
	width: 2rem;
	height: 2rem;
}
/* 录音 */
.rprogress {
	position: absolute;
	left: 50%;
	top: 45%;
	width: 8.75rem;
	height: 8.75rem;
	margin-left: -4.375rem;
	margin-top: -4.375rem;
	background-image: url('data:image/png;base64,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');
	background-repeat: no-repeat;
	background-position: center center;
	background-size: 30px 30px;
	background-color: rgba(0, 0, 0, 0.7);
	border-radius: 0.3125rem;
	transform: 0.15s;
	z-index: 1000;
}
.rschedule {
	background-color: rgba(0, 0, 0, 0);
	border: 0.3125rem solid rgba(0, 183, 229, 0.9);
	opacity: 0.9;
	border-left: 0.3125rem solid rgba(0, 0, 0, 0);
	border-right: 0.3125rem solid rgba(0, 0, 0, 0);
	border-radius: 3.125rem;
	box-shadow: 0 0 0.9375rem #2187e7;
	width: 2.875rem;
	height: 2.875rem;
	position: absolute;
	left: 50%;
	top: 50%;
	margin-left: -1.8rem;
	margin-top: -1.8rem;
	animation: spin 1s infinite linear;
	z-index: 1000;
}
.r-sigh {
	display: none;
	border-radius: 3.125rem;
	box-shadow: 0 0 0.9375rem #2187e7;
	width: 2.875rem;
	height: 2.875rem;
	position: absolute;
	left: 50%;
	top: 50%;
	margin-left: -1.4375rem;
	margin-top: -1.4375rem;
	text-align: center;
	line-height: 2.875rem;
	font-size: 2.5rem;
	font-weight: bold;
	color: #2187e7;
}
.rprogress-sigh {
	background-image: none !important;
}
.rprogress-sigh .rschedule {
	display: none !important;
}
.rprogress-sigh .r-sigh {
	display: block !important;
}
.rsalert {
	font-size: 0.75rem;
	color: #bbb;
	text-align: center;
	position: absolute;
	border-radius: 0.3125rem;
	width: 100%;
	padding: 0.3125rem 0;
	left: 0rem;
	bottom: 0rem;
}
@keyframes spin {
	0% {
		transform: rotate(0deg);
	}
	100% {
		transform: rotate(360deg);
	}
}
body {
	padding: 0;
	margin: 0;
	background-repeat: no-repeat;
	background-attachment: fixed;
}
::-webkit-scrollbar {
	width: 0.625rem;
}
::-webkit-scrollbar-track {
	border-radius: 0.625rem;
	background-color: rgba(25, 147, 147, 0.1);
}
::-webkit-scrollbar-thumb {
	border-radius: 0.625rem;
	background-color: rgba(25, 147, 147, 0.2);
}
.chat-thread {
	margin: 1.5rem auto 0;
	padding: 0 1.5rem 0 0;
	list-style: none;
	overflow-y: scroll;
	overflow-x: hidden;
}
.chat-thread li {
	position: relative;
	clear: both;
	display: inline-block;
	padding: 1rem 2.5rem 1rem 1.25rem;
	margin: 0 0 1.25rem 0;
	font: 1rem/1.25rem 'Noto Sans', sans-serif;
	border-radius: 0.625rem;
	background-color: rgba(25, 147, 147, 0.2);
}
.chat-thread li:before {
	position: absolute;
	top: 0;
	width: 3.125rem;
	height: 3.125rem;
	border-radius: 3.125rem;
	content: '';
}
.chat-thread li:after {
	position: absolute;
	top: 0.9375rem;
	content: '';
	width: 0;
	height: 0;
	border-top: 0.9375rem solid rgba(25, 147, 147, 0.2);
}
.chat-thread li.mine {
	animation: show-chat-odd 0.15s 1 ease-in;
	float: right;
	margin-right: 5rem;
	color: #0ad5c1;
}
.chat-thread li.mine:before {
	right: -5rem;
	background-image: url(data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD/4QBoRXhpZgAATU0AKgAAAAgABAEaAAUAAAABAAAAPgEbAAUAAAABAAAARgEoAAMAAAABAAIAAAExAAIAAAASAAAATgAAAAAAAABgAAAAAQAAAGAAAAABUGFpbnQuTkVUIHYzLjUuMTAA/9sAQwAHBQUGBQQHBgUGCAcHCAoRCwoJCQoVDxAMERgVGhkYFRgXGx4nIRsdJR0XGCIuIiUoKSssKxogLzMvKjInKisq/9sAQwEHCAgKCQoUCwsUKhwYHCoqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioq/8AAEQgAMgAyAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/aAAwDAQACEQMRAD8A8wre0/w55qLLqM62ysMrEWAdh+PSl8M6fFLMbu5K7YziNT3b1/CqniRLq98UA2SlhHGobnA55FdbajHmZwxTnLlRtm78NabDGhs/tEwchmVfMGD0znvn0pqahoN3fCH+z0RcHcWiKY4745/KuUVtTtdSFtJCvmxHmN+317Gp5dTthGZXUSXPJJVsh1PUH1rP2rZt7GKOi1LwxHIv2jRg20jPks4b64b+h/OuZZWRyrqVZTggjBBrZ0PWftFwkA8xImTpu7gVJr1kpX7UhzJ0cZ5I9a0umtDJxcWYVFFFBJ1WlKkOmQASYyoYjZnk81aukuLHTP7YFtDeW6zJEFmyu4555H8PQemaz9NmR9PhO45C7T+HFa1rrk9iwSZ3vLJU+WyaVUCsDncMjJI64H41riYRVHmS7E4KbliOWT7nPePtPvbm8j1trWSyF3GpkgkYEbl4BBHTjHBrn49GkfRLnUpn2tBgtFj76k4GD/e749K63xb4gtdTtp4LeRismGVyOmDnpWBq2pXd74ZsdPjG2GCTdHDEnzSN/ebHLH+VeXBuyR7E4wu35EXhKKObVCRnMKMwY+hwAPz5rrJ7bzYXTfu3Ag5yKwfC+nNZvcPPJHvdQDGnJjOehPTPsOlb0zLFA8m/hVJ/SvXo0oundnhV6slU5UcjRRRXPY3Luk3giYwSHCscqfetfULXOlR3TXIhZ2ZYfLILHHDkjsO3PJPSuXqeC5aPzd5ZjIQdxOcEDFa+0fJyEKmvac5myyQWJlRZnkkTHytjoeuPepLaa9vi4ib7LbOMyFG5KjtnrVMabPd30gcrCjMSZGPQf1NbF0i29pHBAUlbABVT8pPqSOwrljDW53c+lrl/RWxvymyDAWPjsPSpdVnVF8iNsk8tz0HpVFLloowEYvJjBkIwB7AdhUGSxJJyT1JrpU2o8py1IwlJS6hRS0VAiKiiikUKKUUUUxC0ooooAKKKKQH/2Q==);
}
.chat-thread li.mine:after {
	border-right: 0.9375rem solid transparent;
	right: -15px;
}
.chat-thread li.other {
	animation: show-chat-even 0.15s 1 ease-in;
	float: left;
	margin-left: 5rem;
	color: #0ec879;
}
.chat-thread li.other:before {
	left: -5rem;
	background-image: url(data:image/jpeg;base64,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);
}
.chat-thread li.other:after {
	border-left: 0.9375rem solid transparent;
	left: -0.9375rem;
}
.chat-window {
	position: fixed;
	bottom: 1.125rem;
}
.chat-window-message {
	width: 100%;
	height: 3rem;
	font: 32px/48px 'Noto Sans', sans-serif;
	background: none;
	color: #0ad5c1;
	border: 0;
	border-bottom: 0.0625rem solid rgba(25, 147, 147, 0.2);
	outline: none;
}
@media all and (max-width: 767px) {
	.chat-thread {
		width: 100%;
		height: 100%;
	}
	.chat-window {
		left: 5%;
		width: 90%;
	}
}
@media all and (min-width: 768px) {
	.chat-thread {
		width: 100%;
		height: 100%;
	}
	.chat-window {
		left: 25%;
		width: 50%;
	}
}
@keyframes show-chat-even {
	0% {
		margin-left: -30rem;
	}
	100% {
		margin-left: 0;
	}
}
@keyframes show-chat-odd {
	0% {
		margin-right: -30rem;
	}
	100% {
		margin-right: 0;
	}
}
.credits {
	text-align: center;
	margin-top: 2.1875rem;
	color: rgba(255, 255, 255, 0.35);
	font-family: 'Noto Sans', sans-serif;
}
.credits a {
	text-decoration: none;
	color: rgba(255, 255, 255, 0.35);
}
</style>
