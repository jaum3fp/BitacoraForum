<script setup lang="ts">

</script>


<template>

<div class="profile-container">
    <div class="profile-icon">
        <svg class="profile-svg" 
        xmlns="http://www.w3.org/2000/svg"  
        viewBox="0 0 24 24"  
        stroke="currentColor"  
        stroke-width="2"  s
        troke-linecap="round"  
        stroke-linejoin="round">
            <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
            <path d="M8 7a4 4 0 1 0 8 0a4 4 0 0 0 -8 0" />
            <path d="M6 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2" />
        </svg>
    </div>
</div>

</template>


<style scoped>

.profile-container {
    padding: 6vh 0;
}

.profile-icon {
    cursor: pointer;
    width: 5em;
    height: 5em;
    border-radius: 50%;
    background-color: var(--text-color);
    display: flex;
    justify-content: center;
    align-items: center;
    margin: auto;
}

.profile-svg {
    color: var(--primary-color);
    width: 48px;
    fill: none;
}

</style>
