<template>
  <div class="app h-screen w-full overflow-auto overflow-x-auto">
    <div class="grid h-screen min-h-screen w-full overflow-x-auto  lg:grid-cols-[280px_1fr]">
      <div class="hidden border-r bg-gray-100/40 lg:block dark:bg-gray-800/40">
        <div class="flex flex-col gap-2">
          <div class="flex h-[60px] items-center px-6"><a class="flex items-center gap-2 font-semibold"
                                                          href="#">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                 stroke-linejoin="round" class="h-6 w-6">
              <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
              <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
              <path d="M12 3v6"></path>
            </svg>
            <span class="">区块链养老保险平台</span></a></div>
          <div class="flex-1">
            <nav class="grid items-start px-4 text-sm font-medium">
              <router-link to="/"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                  <polyline points="9 22 9 12 15 12 15 22"></polyline>
                </svg>
                主页
              </a></router-link>
              <router-link to="order"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <circle cx="8" cy="21" r="1"></circle>
                  <circle cx="19" cy="21" r="1"></circle>
                  <path
                      d="M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12">
                  </path>
                </svg>
                交易状况
              </a></router-link>
              <router-link to="product"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="m7.5 4.27 9 5.15"></path>
                  <path
                      d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z">
                  </path>
                  <path d="m3.3 7 8.7 5 8.7-5"></path>
                  <path d="M12 22V12"></path>
                </svg>
                产品
              </a></router-link>
              <router-link to="customer"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
                  <circle cx="9" cy="7" r="4"></circle>
                  <path d="M22 21v-2a4 4 0 0 0-3-3.87"></path>
                  <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                </svg>
                顾客
              </a></router-link>
              <router-link to="analytic"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="M3 3v18h18"></path>
                  <path d="m19 9-5 5-4-4-3 3"></path>
                </svg>
                分析
              </a></router-link>
              <router-link to="nodeInfo"><a
                  class="flex items-center gap-3 rounded-lg bg-gray-100 px-3 py-2 text-gray-900  transition-all hover:text-gray-900 dark:bg-gray-800 dark:text-gray-50 dark:hover:text-gray-50"
                  href="#">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="h-4 w-4"
                >
                  <path d="m7.5 4.27 9 5.15"></path>
                  <path
                      d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"></path>
                  <path d="m3.3 7 8.7 5 8.7-5"></path>
                  <path d="M12 22V12"></path>
                </svg>
                节点信息
              </a></router-link>
            </nav>
          </div>
        </div>
      </div>
      <div class="flex flex-col">
        <header class="flex h-14 lg:h-[60px] items-center gap-4 border-b bg-gray-100/40 px-6 dark:bg-gray-800/40"><a
            class="lg:hidden" href="#">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
               viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
               stroke-linejoin="round" class="h-6 w-6">
            <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
            <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
            <path d="M12 3v6"></path>
          </svg>
          <span class="sr-only">Analytic</span></a>
          <div class="flex-1">
            <h1 class="font-semibold text-lg">节点</h1>
          </div>
          <div class="flex flex-1 items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
            <form class="ml-auto flex-1 sm:flex-initial">
            </form>
            <div class="dropdown dropdown-end">
              <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                <div class="w-12 rounded-full">
                  <img alt="Tailwind CSS Navbar component" :src="avatar" style="width: 100%"/>
                </div>
              </div>
              <ul tabindex="0"
                  class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
                <li @click="openVerified">
                  <a class="justify-between">
                    实名认证
                    <span class="badge" v-if="isUserAuthenticated">已认证</span>
                    <span class="badge" v-else>未认证</span>
                  </a>
                </li>
                <li @click="openUpdate"><a>修改用户信息</a></li>
                <li @click="logout"><a>登出</a></li>
                <li @click="userLogout"><a>账号注销</a></li>
              </ul>
            </div>
          </div>
        </header>
        <main class="flex flex-1 flex-col gap-4 overflow-auto p-4 md:gap-8 md:p-6">
          <div class="box" style="width: 100%">
            <div class="max-w-4xl mx-auto my-8">
              <div class="grid grid-cols-4 gap-4 mb-8">
                <div class="rounded-lg border shadow-sm bg-[#29b6f6] text-white" data-v0-t="card">
                  <div class="flex flex-col space-y-1.5 p-6">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="text-white"
                    >
                      <circle cx="12" cy="12" r="10"></circle>
                      <path d="M12 16v-4"></path>
                      <path d="M12 8h.01"></path>
                    </svg>
                    <h3 class="text-2xl font-semibold whitespace-nowrap leading-none tracking-tight">
                      {{ data.node_number }}</h3>
                    <p class="text-sm text-muted-foreground">节点数量</p>
                  </div>
                </div>
                <div class="rounded-lg border shadow-sm bg-[#26c6da] text-white" data-v0-t="card">
                  <div class="flex flex-col space-y-1.5 p-6">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="text-white"
                    >
                      <path d="m3 21 1.9-5.7a8.5 8.5 0 1 1 3.8 3.8z"></path>
                    </svg>
                    <h3 class="text-2xl font-semibold whitespace-nowrap leading-none tracking-tight">
                      {{ parseString(data.transaction_number) }}</h3>
                    <p class="text-sm text-muted-foreground">交易数量</p>
                  </div>
                </div>
                <div class="rounded-lg border shadow-sm bg-[#7e57c2] text-white" data-v0-t="card">
                  <div class="flex flex-col space-y-1.5 p-6">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="text-white"
                    >
                      <circle cx="8" cy="21" r="1"></circle>
                      <circle cx="19" cy="21" r="1"></circle>
                      <path d="M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12"></path>
                    </svg>
                    <h3 class="text-2xl font-semibold whitespace-nowrap leading-none tracking-tight">
                      {{ data.block_number }}</h3>
                    <p class="text-sm text-muted-foreground">区块高度</p>
                  </div>
                </div>
                <div class="rounded-lg border shadow-sm bg-[#ef5350] text-white" data-v0-t="card">
                  <div class="flex flex-col space-y-1.5 p-6">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="text-white"
                    >
                      <path d="M6 4v6a6 6 0 0 0 12 0V4"></path>
                      <line x1="4" x2="20" y1="20" y2="20"></line>
                    </svg>
                    <h3 class="text-2xl font-semibold whitespace-nowrap leading-none tracking-tight">
                      {{ data.pending_transaction_number }}</h3>
                    <p class="text-sm text-muted-foreground">进行中交易</p>
                  </div>
                </div>
              </div>
              <div class="grid grid-cols-3 gap-8" style="overflow-x: auto">
                <div class="col-span-4">
                  <div class="rounded-lg border bg-card text-card-foreground shadow-sm w-full" data-v0-t="card">
                    <div class="flex flex-col space-y-1.5 p-6">
                      <h3 class="text-2xl font-semibold whitespace-nowrap leading-none tracking-tight">节点信息</h3>
                    </div>
                    <div class="p-6 overflow-x-auto">
                      <div
                          dir="ltr"
                          class="relative w-full border rounded-md h-72"
                          style="position: relative; --radix-scroll-area-corner-width: 0px; --radix-scroll-area-corner-height: 0px;"
                      >
                        <div
                            data-radix-scroll-area-viewport=""
                            class="h-full w-full rounded-[inherit] overflow-x-auto"
                        >
                          <div style="min-width: 100%; display: table;">
                            <div class="relative w-full overflow-auto">
                              <table class="w-full caption-bottom text-sm">
                                <thead class="[&amp;_tr]:border-b">
                                <tr class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                                  <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                                    ID
                                  </th>
                                  <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                                    区块高度
                                  </th>
                                  <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                                    节点ID
                                  </th>
                                  <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                                    状态
                                  </th>
                                </tr>
                                </thead>
                                <tbody class="[&amp;_tr:last-child]:border-0">
                                <tr class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted"
                                    v-for="(k,v) in data.node_info_list" :key="k">
                                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 font-medium">
                                    {{ v }}
                                  </td>
                                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">
                                    {{ data.block_number }}
                                  </td>
                                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">{{ k }}</td>
                                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">
                                    <div
                                        class="inline-flex items-center rounded-full whitespace-nowrap border px-2.5 py-0.5 w-fit text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80">
                                      运行
                                    </div>
                                  </td>
                                </tr>
                                </tbody>
                              </table>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "nodeInfo",
  data() {
    return {
      avatar: '',
      isUserAuthenticated: false,
      data: {}
    }
  },
  mounted() {
  },
  created() {
    const authValue = localStorage.getItem('user_auth');
    this.isUserAuthenticated = authValue === 'true'; // 确保将字符串 'true' 转换为布尔值 true
    this.avatar = localStorage.getItem("avatar");
    this.$axios.get("/main/getNodeInfo", {
      'Authorization': `Bearer ` + localStorage.getItem('auth_token')
    }).then((res) => {
      if (res.data.code == 200) {
        this.data = res.data.data;
      }
      console.log(res)
    }).catch((res) => {
      console.log(res)
    })
  },
  methods: {
    // 账号注销接口
    userLogout() {
      this.$axios.post('/user/logout', {}, {
        headers: {
          'Authorization': `Bearer ` + localStorage.getItem('auth_token')
        }
      }).then((res) => {
        console.log(res)
      }).catch((error) => {
        console.error(error)
      })

      // 删除所有缓存
      localStorage.removeItem("auth_token")
      localStorage.removeItem("user_auth")
      // 将路由切换到 /login
      this.$router.push('/login')
    },
    parseString(hexString) {
      return parseInt(hexString, 16);
    },
    // 账号登出
    logout() {
      // 删除所有缓存
      localStorage.removeItem("auth_token")
      localStorage.removeItem("user_auth")
      // 将路由切换到 /login
      this.$router.push('/login')
    },
    openVerified() {
      if (this.isUserAuthenticated == false) {
        window.location.href = '#my_modal_8';
      }
    },
    openUpdate() {
      if (this.isUserAuthenticated == true) {
        window.location.href = '#updateUserInfo';
      }
    },
    getRandomColor() {
      const letters = '0123456789ABCDEF';
      let color = '#';
      for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    },
    getAvatarStyle(user) {
      // 尝试从 localStorage 中获取颜色
      let colorStart = localStorage.getItem("colorStart");
      let colorEnd = localStorage.getItem("colorEnd");

      // 如果没有找到缓存的颜色，生成新的颜色
      if (!colorStart || !colorEnd) {
        colorStart = this.getRandomColor();
        colorEnd = this.getRandomColor();

        // 将生成的颜色存入 localStorage
        localStorage.setItem("colorStart", colorStart);
        localStorage.setItem("colorEnd", colorEnd);
      }

      // 通过指定色标的位置来增加渐变尺度
      // 比如从 10% 的位置开始到 90% 的位置结束
      return {
        background: `linear-gradient(to right, ${colorStart}, ${colorEnd})`,
        width: '32px',
        height: '32px',
        borderRadius: '50%',
        aspectRatio: '1 / 1',
        objectFit: 'cover'
      };
    },
  }

}
</script>

<style scoped>
.avatar {
  display: inline-block;
  overflow: hidden;
  /* 如果你要在里面放图标或者文字，建议设置 overflow */
  /* 其他样式 */
}

[data-radix-scroll-area-viewport] {
  scrollbar-width: none;
  -ms-overflow-style: none;
  -webkit-overflow-scrolling: touch;
}

[data-radix-scroll-area-viewport]::-webkit-scrollbar {
  display: none;
}

[data-radix-scroll-area-viewport] {
  scrollbar-width: none;
  -ms-overflow-style: none;
  -webkit-overflow-scrolling: touch;
}

[data-radix-scroll-area-viewport]::-webkit-scrollbar {
  display: none
}
</style>