var Util = {
	/**
	 *校验手机号
	 * @param tel
	 * @returns {boolean}
	 */
	isMobile: function(tel) {
		let pattern = /^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^17[0,6,7,8]{1}\d{8}$|^18[\d]{9}$/;
		return pattern.test(tel)
	},
	/**
	 * 校验密码
	 * @param password
	 * @returns {boolean}
	 */
	isPassword:function(pwd){
		let pattern = /((?=.*[a-z])(?=.*\d)|(?=[a-z])(?=.*[#@!~%^&*._$,?+=-])|(?=.*\d)(?=.*[#@!~%^&*._$,?+=-]))[a-z\d#@!~%^&*._$,?+=-]{6,16}/i
		return pattern.test(pwd)
	},
	/**
	 * 校验邮箱
	 * @param email
	 * @constructor
	 */
	isEmail: function(email) {
		let pattern = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/
		return pattern.test(email);
	},
	/**
	 * 校验身份证
	 * @param idCard
	 * @constructor
	 */
	isIDCard: function(idCard) {
		let pattern =
			/^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$|^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}([0-9]|X)$/;
		return pattern.test(idCard);
	},
	/**
	 * 数字格式化加逗号
	 * @param num
	 * @returns {string}
	 */
	toThousands: function(num) {
		num = (num || 0).toString();
		let result = '';
		while (num.length > 3) {
			result = ',' + num.slice(-3) + result;
			num = num.slice(0, num.length - 3);
		}
		if (num) {
			result = num + result;
		}
		return result;
	},
}
export default Util
