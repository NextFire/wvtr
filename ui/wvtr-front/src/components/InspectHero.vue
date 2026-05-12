<script setup lang="ts">
    import { inject } from 'vue'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'
    import { clampToRange } from '@/tools/utils.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const hero = navigationHandler.getHeroToInspect()

    function getClampedHP(currentHP: number, maxHP: number) {
        return clampToRange(currentHP, 0, maxHP)
    }
</script>

<template>
    <section v-if="hero" class="inspect-sheet">
        <div class="panel panel-feature inspect-hero-card">
            <div class="inspect-hero-grid">
                <div class="inspect-hero-media">
                    <div>
                        <p class="eyebrow">Hero dossier</p>
                        <h2>{{ hero.name }}</h2>
                        <p>{{ hero.heroClass.description }}</p>
                    </div>

                    <div class="card-meta-row">
                        <span class="badge-pill">{{ hero.heroClass.name }}</span>
                        <span class="badge-pill badge-pill-soft">Rank {{ hero.rank }}</span>
                        <span class="badge-pill badge-pill-soft">Level {{ hero.attributes.level }}</span>
                    </div>

                    <div class="waifu-image-container inspect-portrait">
                        <img :src="hero.imageUrl" :alt="hero.name"/>
                    </div>
                </div>

                <div class="inspect-hero-body">
                    <section class="panel inset-panel">
                        <p class="eyebrow">Vital bars</p>
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
                    </section>

                    <section class="stat-grid">
                        <div class="stat-tile">
                            <span>Strength</span>
                            <strong>{{ hero.attributes.strength }}</strong>
                            <p>Growth {{ hero.attributes.sgt.toFixed(2) }}</p>
                        </div>
                        <div class="stat-tile">
                            <span>Intelligence</span>
                            <strong>{{ hero.attributes.intelligence }}</strong>
                            <p>Growth {{ hero.attributes.igt.toFixed(2) }}</p>
                        </div>
                        <div class="stat-tile">
                            <span>Dexterity</span>
                            <strong>{{ hero.attributes.dexterity }}</strong>
                            <p>Growth {{ hero.attributes.dgt.toFixed(2) }}</p>
                        </div>
                        <div class="stat-tile">
                            <span>Luck</span>
                            <strong>{{ hero.attributes.luck }}</strong>
                            <p>Growth {{ hero.attributes.lgt.toFixed(2) }}</p>
                        </div>
                    </section>

                    <section class="panel inset-panel">
                        <p class="eyebrow">Resistances</p>
                        <div class="resistance-grid">
                            <div class="stat-tile">
                                <span>Blunt</span>
                                <strong>{{ hero.attributes.blunt }}</strong>
                            </div>
                            <div class="stat-tile">
                                <span>Pierce</span>
                                <strong>{{ hero.attributes.pierce }}</strong>
                            </div>
                            <div class="stat-tile">
                                <span>Slash</span>
                                <strong>{{ hero.attributes.slash }}</strong>
                            </div>
                            <div class="stat-tile">
                                <span>Fire</span>
                                <strong>{{ hero.attributes.fire }}</strong>
                            </div>
                            <div class="stat-tile">
                                <span>Frost</span>
                                <strong>{{ hero.attributes.frost }}</strong>
                            </div>
                            <div class="stat-tile">
                                <span>Lightning</span>
                                <strong>{{ hero.attributes.lighting }}</strong>
                            </div>
                        </div>
                    </section>

                    <section class="panel inset-panel">
                        <p class="eyebrow">Skill loadout</p>
                        <div class="skill-grid">
                            <article class="skill-card">
                                <template v-if="hero.weaponAttack">
                                    <img v-if="hero.weaponAttack.image_url !== ''" class="skill-icon" :src="hero.weaponAttack.image_url" :alt="hero.weaponAttack.name"/>
                                    <h3>{{ hero.weaponAttack.name }}</h3>
                                    <p>{{ hero.weaponAttack.description }}</p>
                                </template>
                                <template v-else>
                                    <h3>No weapon skill</h3>
                                    <p>This hero does not have a weapon skill assigned yet.</p>
                                </template>
                            </article>

                            <article class="skill-card">
                                <template v-if="hero.uniqueSkill">
                                    <img v-if="hero.uniqueSkill.image_url !== ''" class="skill-icon" :src="hero.uniqueSkill.image_url" :alt="hero.uniqueSkill.name"/>
                                    <h3>{{ hero.uniqueSkill.name }}</h3>
                                    <p>{{ hero.uniqueSkill.description }}</p>
                                </template>
                                <template v-else>
                                    <h3>No unique skill</h3>
                                    <p>This hero has not received a unique skill yet.</p>
                                </template>
                            </article>

                            <article class="skill-card">
                                <template v-if="hero.activeSkill">
                                    <img v-if="hero.activeSkill.image_url !== ''" class="skill-icon" :src="hero.activeSkill.image_url" :alt="hero.activeSkill.name"/>
                                    <h3>{{ hero.activeSkill.name }}</h3>
                                    <p>{{ hero.activeSkill.description }}</p>
                                </template>
                                <template v-else>
                                    <h3>No active skill</h3>
                                    <p>This hero does not have an active skill configured yet.</p>
                                </template>
                            </article>
                        </div>
                    </section>
                </div>
            </div>
        </div>
    </section>

    <section v-else class="inspect-sheet">
        <div class="panel">
            <p class="eyebrow">Hero dossier</p>
            <h2>No hero selected</h2>
            <p>Pick a portrait anywhere in the guild hall to inspect a hero in detail.</p>
        </div>
    </section>
</template>
