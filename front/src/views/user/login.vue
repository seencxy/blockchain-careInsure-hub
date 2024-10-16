<template>
  <div class="app">
    <div class="success" style="position: absolute; right:5%;top:3%" v-show="alert">
      <div role="alert" class="alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none"
             viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <span v-text="message"></span>
      </div>
    </div>
    <div class="success" style="position: absolute; right:5%;top:3%" v-show="success">
      <div role="alert" class="alert alert-success">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none"
             viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <span v-text="message"></span>
      </div>
    </div>
    <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <div class="w-full max-w-xs p-8 bg-white rounded-2xl shadow-md">
        <div class="mb-4 text-center"><span class="text-5xl font-bold text-blue-500">HNKJZYXY</span></div>
        <h2 class="mb-8 text-2xl font-semibold text-center text-gray-700">区块链养老保险平台</h2>
        <form class="flex flex-col space-y-4"><input
            class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            id="text" placeholder="用户名" v-model="this.username"><input
            class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            id="password" placeholder="密码" type="password" v-model="this.password"><a
            class="text-sm text-blue-600 hover:underline" @click="gotoPasswordReset"
            style="cursor: pointer;">忘记密码?</a>
          <button
              class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-4 py-2 mt-4 bg-blue-500 text-white"
              @click="login">登录
          </button>
        </form>
        <p class="mt-4 text-center text-sm">
          还没有账号? <a class="text-blue-600 hover:underline" style="cursor: pointer;" @click="gotoRegister">注册</a>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      username: '',
      password: '',
      alert: false,
      success: false,
      message: ''
    }
  },
  methods: {
    // 跳转到找回密码页面
    gotoPasswordReset() {
      this.$router.push('/password_reset')
    },
    // 跳转到注册页面
    gotoRegister() {
      this.$router.push('/register')
    },
    // 登录接口
    login(event) {
      // 阻止点击默认事件发生
      event.preventDefault();
      if (this.username.length < 6 || this.password.length < 6) { // 假设用户名和密码长度都应大于等于6
        this.message = '用户名或密码错误'
        this.alert = true;
        setTimeout(() => {
          this.alert = false;
        }, 1000); // 1秒后执行
      } else {
        this.$axios.post('/user/login', {
          username: this.username,
          password: this.password
        }).then((rep) => {
          console.log(rep)
          if (rep.data.code != 200) {
            this.message = rep.data.message
            this.alert = true;
            this.password = ''
            setTimeout(() => {
              this.alert = false;
            }, 1000); // 1秒后执行
            return
          } else {
            this.message = rep.data.message
            // 设置访问令牌
            localStorage.setItem("auth_token", rep.data.token)
            // 判断用户登录账号是否认证
            localStorage.setItem("user_auth", rep.data.flag)
            // 设置用户头像
            localStorage.setItem("avatar", rep.data.avatar)
            this.success = true
            setTimeout(() => {
              this.success = false;
               // 跳转到首页
            if (this.username === "admin123") {
              this.$router.push('/admin/home')
            } else {
              this.$router.push('/')
            }
            }, 1000); // 1秒后执行
            
          }
        }).catch((rep) => {
          this.message = rep.message
          this.alert = true;
          this.username = ''
          this.password = ''
          setTimeout(() => {
            this.alert = false;
          }, 1000); // 1秒后执行
        })
      }
    }
  }
}
</script>

<style scoped></style>