<script setup lang="ts">
    import { inject } from "vue";
import { EquipmentType, type Armor, type Hero, type Omamori, type Weapon } from "../tools/types.ts"
    import { global } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts";

    const props = defineProps<{
        weapon: Weapon | undefined;
        armor: Armor | undefined;
        omamori: Omamori | undefined;
    }>();

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const eqType = navigationHandler.getInventoryType()
    
</script>

<template>
    <div v-if="eqType == EquipmentType.WeaponType" class="inspect-hero">
        <div> 
            <h1>{{ weapon?.name }}</h1>
            <img :src="weapon?.iconURL" width="150px">
        </div>
        <div>damage: {{ weapon?.baseDamage }}</div>
        <div>crit rate: {{ weapon?.baseCritRate }}</div>
        <div>attack speed: {{ weapon?.baseAttackSpeed }}</div>
        <div>
            <div v-for="a in weapon?.affixes">
                <div>effect : {{ a.type }}</div>
            </div>
        </div>
    </div>
    <div v-else-if="eqType == EquipmentType.ArmorType" class="inspect-hero">
        <div> 
            <h1>{{ armor?.name }}</h1>
            <img :src="armor?.iconURL" width="150px">
        </div>
        <div>resistances: {{ armor?.baseResistancesRange }}</div>
        <div>block: {{ armor?.blockScore }}</div>
        <div>evade: {{ armor?.evadeScore }}</div>
        <div>
            <div v-for="a in armor?.affixes">
                <div>effect : {{ a.type }}</div>
            </div>
        </div>
    </div>
    <div v-else-if="eqType == EquipmentType.OmamoriType" class="inspect-hero">
        <div> 
            <h1>{{ omamori?.name }}</h1>
            <img :src="omamori?.iconURL" width="150px">
        </div>
        <div>
            <div v-for="a in omamori?.affixes">
                <div>effect : {{ a.type }}</div>
            </div>
        </div>
    </div>
    <div v-else>
        no equipment to inspect
    </div>
</template>