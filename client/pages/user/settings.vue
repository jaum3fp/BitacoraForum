<script setup lang="ts">
import ProfileImageField from '~/components/Forms/Fields/ProfileImageField.vue';
import { UserModel, type UserModelType } from '~/models/user';

</script>

<template>
<div class="rules p-8 m-8 text-left rounded-lg bg-white/10 text-[var(--text-color)]">
    <div>
        <h1>Settings</h1>
        <I18nSelect class="w-64" />
    </div>
</div>
</template>
