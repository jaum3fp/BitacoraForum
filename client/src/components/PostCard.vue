<script setup lang="ts">

defineProps<{
    title: string,
    description: string,
    author: string,
    image: string,
}>()

</script>


<template>

<div class="post-card">
    <div class="post-card-header">
    <img class="post-card-img" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fas2.ftcdn.net%2Fv2%2Fjpg%2F05%2F38%2F01%2F87%2F1000_F_538018794_tp2FHTwEXbtE26z7F0mVrKgXJLD7GzvM.jpg&f=1&nofb=1&ipt=f9c82416a91725bfc917eabec47538591a3800ebe84480886005a8dd37beb795&ipo=images"
        alt="">
    </div>
    <div class="post-body">
    <h2>{{ title }}</h2>
    <p>{{ description }}</p>
    </div>
    <div class="post-author">
    <p>{{ author }}</p>
    </div>
</div>

</template>


<style scoped>

.post-card {
    height: 15vh;
    border: 1px solid #f2f2f2;
    border-radius: 3px;

    display: flex;
    gap: 1em;
}

.post-card-header {
    max-width: 30vh;
    overflow: hidden;
}

.post-card-body {
    flex: 1;
    text-align: left;

    display: flex;
    flex-direction: column;
    gap: 1em;
}

.post-card-img {
    width: 8vw;
}

</style>
