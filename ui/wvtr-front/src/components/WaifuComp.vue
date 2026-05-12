<script setup lang="ts">
    import type { Waifu } from '../tools/types.ts'

    const props = defineProps<{
        waifu: Waifu;
    }>()

    let rankTextClass = 'waifu-rank-e'

    switch (props.waifu.rank) {
        case 'S':
            rankTextClass = 'waifu-rank-s'
            break
        case 'A':
            rankTextClass = 'waifu-rank-a'
            break
        case 'B':
            rankTextClass = 'waifu-rank-b'
            break
        case 'C':
            rankTextClass = 'waifu-rank-c'
            break
        case 'D':
            rankTextClass = 'waifu-rank-d'
            break
        case 'E':
            rankTextClass = 'waifu-rank-e'
            break
    }
</script>

<template>
    <article class="waifu-card">
        <div class="waifu-card-header">
            <p class="eyebrow">Rank {{ props.waifu.rank }}</p>
            <h3>{{ props.waifu.name_user_preferred }}</h3>
        </div>

        <div class="waifu-image-container waifu-art">
            <img :src="props.waifu.image_large" :alt="props.waifu.name_user_preferred"/>
            <div :class="rankTextClass">{{ props.waifu.rank }}</div>
        </div>
    </article>
</template>
