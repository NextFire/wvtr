<script setup lang="ts">
    import { inject } from "vue";
import type { Hero } from "../tools/types.ts"
    import { global } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const hero = navigationHandler.getHeroToInspect()
    console.log(hero.value)
</script>

<template>
    <div v-if="hero" class="inspect-hero">
        <h1>{{ hero.name }}<label style="font-size: medium;"> lvl.{{ hero.attributes.level }}</label></h1>
        <div class="row">
            <div>
                <img :src="hero.imageUrl" width="200"/>
            </div>
            <div class="column">
                <div>
                    <div>Class: {{ hero.heroClass.name }}</div>
                </div>
                <div>
                    <div class="column" style="display: inline-flex;">
                    <label>hp : {{ hero.attributes.currentHP.toFixed(0) }}/{{ hero.attributes.maxHP.toFixed(0) }} </label>
                    <progress :max="hero.attributes.maxHP.toFixed(0)" :value="hero.attributes.currentHP.toFixed(0)" class="hero-progress"/>
                    <label>xp : {{ hero.attributes.currentXP.toFixed(0) }}/{{ hero.attributes.xpBeforLvlUp.toFixed(0) }} </label>
                    <progress :max="hero.attributes.xpBeforLvlUp.toFixed(0)" :value="hero.attributes.currentXP.toFixed(0)" class="hero-progress"/>
                </div>
                </div>
                <div>
                    <label>Attributes:</label>
                    <div>Strength ({{ hero.attributes.sgt.toFixed(2) }}) : {{ hero.attributes.strength }}</div>
                    <div>Intelligence ({{ hero.attributes.igt.toFixed(2) }}) : {{ hero.attributes.intelligence }}</div>
                    <div>Dexterity ({{ hero.attributes.dgt.toFixed(2) }}) : {{ hero.attributes.dexterity }}</div>
                    <div>Luck ({{ hero.attributes.lgt.toFixed(2) }}) : {{ hero.attributes.luck }}</div>
                </div>
                <div>
                    <div>Resitances :</div>
                    <div class="row" style="gap: 10px;">
                        <div>
                            <div>Blunt:</div>
                            <div>Pierce:</div>
                            <div>Slash:</div>
                        </div>
                        <div>
                            <div>{{ hero.attributes.blunt }}</div>
                            <div>{{ hero.attributes.pierce }}</div>
                            <div>{{ hero.attributes.slash }}</div>
                        </div>
                        <div>
                            <div>Fire:</div>
                            <div>Frost:</div>
                            <div>Lightning:</div>
                        </div>
                        <div>
                            <div>{{ hero.attributes.fire }}</div>
                            <div>{{ hero.attributes.frost }}</div>
                            <div>{{ hero.attributes.lighting }}</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="column">
                <label>Skills</label>
                <div class="raw">
                    <div v-if="hero.uniqueSkill" class="column">
                        <label>{{ hero.uniqueSkill.name }}</label>
                        <img v-if="hero.uniqueSkill.image_url !== ''" :src="global.DOMAIN_NAME + hero.uniqueSkill.image_url"/>
                    </div>
                    <div v-if="hero.activeSkill" class="column">
                        <label>{{ hero.activeSkill.name }}</label>
                        <img v-if="hero.activeSkill.image_url !== ''" :src="global.DOMAIN_NAME + hero.activeSkill.image_url"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        no hero to inspect
    </div>
</template>