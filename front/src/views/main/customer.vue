<template>
  <div class="app">
    <div class="grid h-screen min-h-screen w-full overflow-hidden lg:grid-cols-[280px_1fr]">
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
                  class="flex items-center gap-3 rounded-lg bg-gray-100 px-3 py-2 text-gray-900  transition-all hover:text-gray-900 dark:bg-gray-800 dark:text-gray-50 dark:hover:text-gray-50"
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
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
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
          <span class="sr-only">Customer</span></a>
          <div class="flex-1">
            <h1 class="font-semibold text-lg">顾客</h1>
          </div>
          <div class="flex flex-1 items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
            <form class="ml-auto flex-1 sm:flex-initial" @submit.prevent="handleSubmit">
              <div class="relative">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round"
                     class="absolute left-2.5 top-2.5 h-4 w-4 text-gray-500 dark:text-gray-400">
                  <circle cx="11" cy="11" r="8"></circle>
                  <path d="m21 21-4.3-4.3"></path>
                </svg>
                <input
                    class="flex h-10 w-full rounded-md border border-input px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px] bg-white"
                    placeholder="搜索用户" type="search" v-model="search" @keyup.enter="searchFc"></div>
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
        <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">
          <div class="border shadow-sm rounded-lg p-2">
            <div class="relative w-full overflow-auto">
              <table class="w-full caption-bottom text-sm">
                <thead class="[&amp;_tr]:border-b">
                <tr class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 w-[100px]">
                    用户ID
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[150px]">
                    用户名
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    性别
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                    年龄
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                    套餐种类
                  </th>
                  <th
                      class="h-12 px-4 align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-left">
                    入保时间
                  </th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="order in orders" :key="order.id"
                    class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                  <td class="p-4 align-middle [&:has([role=checkbox])]:pr-0 font-medium">{{
                      order.id
                    }}
                  </td>
                  <td class="p-4 align-middle [&:has([role=checkbox])]:pr-0">{{ order.username }}</td>
                  <td class="p-4 align-middle [&:has([role=checkbox])]:pr-0 hidden md:table-cell">{{
                      order.gender
                    }}
                  </td>
                  <td class="p-4 align-middle [&:has([role=checkbox])]:pr-0">{{ order.age }}</td>
                  <td class="p-4 align-middle [&:has([role=checkbox])]:pr-0 text-right">{{
                      order.kind
                    }}
                  </td>
                  <td class="p-4 align-middle [&:has([role=checkbox])]:pr-0">{{ order.timestamp }}
                  </td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="join" style="margin-left: 40%; margin-right: 40%;">
            <button class="join-item btn" @click="prevPage">«</button>
            <button class="join-item btn">Page {{ page }}</button>
            <button class="join-item btn" @click="nextPage">»</button>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "customer",
  data() {
    return {
      isUserAuthenticated: false,
      search: '',
      page: 1,
      orders: [],
      avatar: '',
    }
  },
  created() {
    this.avatar = localStorage.getItem("avatar");
    const authValue = localStorage.getItem('user_auth');
    this.isUserAuthenticated = authValue === 'true'; // 确保将字符串 'true' 转换为布尔值 true
    this.$axios.get("/main/getBuyerInfo", {
      params: {
        id: this.page,
        username: this.search,
      }
    }).then((rep) => {
      console.log(rep);
      this.orders = rep.data.data;
    }).catch((rep) => {
      console.log(rep);
    });

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
    searchFc() {
      this.page = 1,
          this.$axios.get("/main/getBuyerInfo", {
            params: {
              id: this.page,
              username: this.search,
            }
          }).then((rep) => {
            console.log(rep);
            this.orders = rep.data.data;
          }).catch((rep) => {
            console.log(rep);
          });
    },
    prevPage() {
      if (this.page > 1) {
        this.page -= 1;
        this.$axios.get("/main/getBuyerInfo", {
          params: {
            id: this.page,
            username: this.search,
          }
        }).then((rep) => {
          console.log(rep);
          this.orders = rep.data.data;
        }).catch((rep) => {
          console.log(rep);
        });
      }
    },
    nextPage() {
      // 这里您可能需要一个逻辑来检查是否已经是最后一页
      if (this.orders.length == 10) {
        this.page += 1;
      } else {
        return
      }
      this.$axios.get("/main/getBuyerInfo", {
        params: {
          id: this.page,
          username: this.search,
        }
      }).then((rep) => {
        console.log(rep);
        this.orders = rep.data.data;
      }).catch((rep) => {
        console.log(rep);
      });
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

};
</script>

<style scoped>
.avatar {
  display: inline-block;
  overflow: hidden;
  /* 如果你要在里面放图标或者文字，建议设置 overflow */
  /* 其他样式 */
}
</style>