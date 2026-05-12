<script setup lang="ts">
    import { computed, inject } from 'vue'
    import { EncounterState, type Hero } from '../tools/types.ts'
    import { clampToRange, global } from '../tools/utils.ts'
    import InspectButton from './InspectButton.vue'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'

    defineProps<{
        hero: Hero | undefined;
    }>()

    const navigationHandler = inject<NavigationHandler>('navigationHandler')
    const canInspectHero = computed(() => navigationHandler?.getUserState() === EncounterState.Home)

    function getClampedHP(currentHP: number, maxHP: number) {
        return clampToRange(currentHP, 0, maxHP)
    }
</script>

<template>
    <article v-if="hero" class="hero hero-card">
        <div class="waifu-image-container hero-portrait">
            <img :src="hero.imageUrl" :alt="hero.name"/>
            <InspectButton v-if="canInspectHero" :hero="hero"/>
        </div>

        <div class="column hero-copy">
            <div class="hero-heading">
                <div>
                    <p class="eyebrow">Level {{ hero.attributes.level }}</p>
                    <h3>{{ hero.name }}</h3>
                </div>
                <span class="badge-pill">{{ hero.heroClass?.name ?? 'Enemy' }}</span>
            </div>

            <div class="metric-stack">
                <div class="metric-row">
                    <span>HP</span>
                    <strong>{{ getClampedHP(hero.attributes.currentHP, hero.attributes.maxHP).toFixed(0) }}/{{ hero.attributes.maxHP.toFixed(0) }}</strong>
                </div>
                <progress :max="hero.attributes.maxHP" :value="getClampedHP(hero.attributes.currentHP, hero.attributes.maxHP)" class="hero-progress hero-progress-health"/>

                <div class="metric-row">
                    <span>XP</span>
                    <strong>{{ hero.attributes.currentXP.toFixed(0) }}/{{ hero.attributes.xpBeforLvlUp.toFixed(0) }}</strong>
                </div>
                <progress :max="hero.attributes.xpBeforLvlUp.toFixed(0)" :value="hero.attributes.currentXP.toFixed(0)" class="hero-progress hero-progress-xp"/>
            </div>
        </div>
    </article>

    <article v-else class="hero hero-card hero-card-empty">
        <div class="waifu-image-container hero-portrait">
            <img :src="global.NO_IMAGE" alt="Empty party slot"/>
        </div>

        <div class="column hero-copy">
            <div class="hero-heading">
                <div>
                    <p class="eyebrow">Vacant slot</p>
                    <h3>Open Position</h3>
                </div>
                <span class="badge-pill badge-pill-muted">Recruit me</span>
            </div>

            <p class="hero-empty-copy">This spot is waiting for the next hero you forge, recruit, or lovingly overbuild.</p>

            <div class="metric-stack">
                <div class="metric-row">
                    <span>HP</span>
                    <strong>0/0</strong>
                </div>
                <progress :max="100" :value="0" class="hero-progress hero-progress-health"/>

                <div class="metric-row">
                    <span>XP</span>
                    <strong>0/0</strong>
                </div>
                <progress :max="100" :value="0" class="hero-progress hero-progress-xp"/>
            </div>
        </div>
    </article>
</template>
