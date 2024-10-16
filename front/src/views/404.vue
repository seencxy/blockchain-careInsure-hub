<template>
    <div class="app">
        <section class="w-full h-screen flex flex-col items-center justify-center px-4 text-center bg-gray-100">
            <div class="space-y-4">
                <h1 class="text-6xl font-bold text-gray-900">
                    404
                </h1>
                <p class="mx-auto max-w-[700px] text-xl text-gray-600">
                    哎呀！您要查找的页面不存在。
                </p>
            </div><a
                class="inline-flex h-10 items-center justify-center rounded-md bg-gray-900 px-8 text-sm font-medium text-gray-50 shadow transition-colors hover:bg-gray-900/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-950 disabled:pointer-events-none disabled:opacity-50 dark:bg-gray-50 dark:text-gray-900 dark:hover:bg-gray-50/90 dark:focus-visible:ring-gray-300"
                @click="backHome" href="/">
                回到主页
            </a>
        </section>
    </div>
</template>

<script>
export default {
    name: "404",
    methods: {
        // 回到主页
        backHome() {
            // 跳转到主页
            this.$router.push('/');
        }
    }
}
</script>

<style scoped></style>