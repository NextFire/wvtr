<script setup lang="ts">
    import { computed, inject } from 'vue'
    import Team from './Team.vue'
    import TeamManagement from './TeamManagement.vue'
    import ExpeditionsList from './ExpeditionsList.vue'
    import Waifus from './Waifus.vue'
    import InspectHero from './InspectHero.vue'
    import { NavigationStatus, NavigationHandler } from '@/tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
    const navigationStatus = navigationHandler.getNavigationStatus()

    const activeModeLabel = computed(() => {
        switch (navigationStatus.value) {
            case NavigationStatus.TeamManagement:
                return 'Team Forge'
            case NavigationStatus.ExpeditionManagement:
                return 'Expedition Board'
            case NavigationStatus.HeroMaker:
                return 'Hero Summoning'
            case NavigationStatus.InspectHero:
                return 'Hero Dossier'
            default:
                return 'Guild Hall'
        }
    })

    function setHomeStatus(newStatus: NavigationStatus) {
        navigationHandler.setHomeStatus(newStatus)
    }
</script>

<template>
    <section class="home-screen">
        <section class="panel panel-feature home-hero-banner">
            <div class="home-banner-copy">
                <p class="eyebrow">Guild hall</p>
                <h2>Build your squad, recruit heroes, and launch the next adventure.</h2>
                <p>
                    Your roster, expeditions, and hero tools now live on one playful command board.
                </p>
            </div>
            <div class="home-banner-stats">
                <div class="stat-pill">
                    <span>Roster Size</span>
                    <strong>{{ user?.ownedHeroes.length ?? 0 }}</strong>
                </div>
                <div class="stat-pill">
                    <span>Active Party</span>
                    <strong>{{ user?.currentTeam?.heroes.length ?? 0 }}/3</strong>
                </div>
                <div class="stat-pill">
                    <span>Current Mode</span>
                    <strong>{{ activeModeLabel }}</strong>
                </div>
            </div>
        </section>

        <div class="home-grid">
            <aside class="panel home-team-panel">
                <div class="panel-heading">
                    <div>
                        <p class="eyebrow">Current party</p>
                        <h3>Frontline Squad</h3>
                    </div>
                    <button class="ghost-button" v-on:click="setHomeStatus(NavigationStatus.TeamManagement)">Edit lineup</button>
                </div>
                <Team :team="user?.currentTeam"/>
            </aside>

            <section class="home-content">
                <div class="home-nav">
                    <button
                        class="tab-button"
                        :class="{ 'is-active': navigationStatus == NavigationStatus.Home }"
                        v-on:click="setHomeStatus(NavigationStatus.Home)"
                    >
                        Overview
                    </button>
                    <button
                        class="tab-button"
                        :class="{ 'is-active': navigationStatus == NavigationStatus.TeamManagement }"
                        v-on:click="setHomeStatus(NavigationStatus.TeamManagement)"
                    >
                        Team Forge
                    </button>
                    <button
                        class="tab-button"
                        :class="{ 'is-active': navigationStatus == NavigationStatus.ExpeditionManagement }"
                        v-on:click="setHomeStatus(NavigationStatus.ExpeditionManagement)"
                    >
                        Expeditions
                    </button>
                    <button
                        class="tab-button"
                        :class="{ 'is-active': navigationStatus == NavigationStatus.HeroMaker }"
                        v-on:click="setHomeStatus(NavigationStatus.HeroMaker)"
                    >
                        Summon Heroes
                    </button>
                    <button
                        v-if="navigationStatus == NavigationStatus.InspectHero"
                        class="tab-button is-active"
                    >
                        Dossier
                    </button>
                </div>

                <section v-if="navigationStatus == NavigationStatus.Home" class="panel home-overview-panel">
                    <div class="panel-heading">
                        <div>
                            <p class="eyebrow">Quick actions</p>
                            <h3>Choose your next bit of trouble</h3>
                        </div>
                    </div>

                    <div class="home-overview-cards">
                        <button class="action-card" v-on:click="setHomeStatus(NavigationStatus.TeamManagement)">
                            <p class="eyebrow">Squad Studio</p>
                            <h3>Shape your frontline</h3>
                            <p>Swap heroes, test combinations, and lock in your next team build.</p>
                        </button>

                        <button class="action-card" v-on:click="setHomeStatus(NavigationStatus.ExpeditionManagement)">
                            <p class="eyebrow">Quest Counter</p>
                            <h3>Launch a new expedition</h3>
                            <p>Browse mission cards, compare duration and cost, and commit your party.</p>
                        </button>

                        <button class="action-card" v-on:click="setHomeStatus(NavigationStatus.HeroMaker)">
                            <p class="eyebrow">Summoning Room</p>
                            <h3>Turn waifus into heroes</h3>
                            <p>Pick a favorite, promote them into the roster, and inspect the result instantly.</p>
                        </button>
                    </div>
                </section>

                <section v-else class="panel home-panel">
                    <TeamManagement v-if="navigationStatus == NavigationStatus.TeamManagement"/>
                    <ExpeditionsList v-else-if="navigationStatus == NavigationStatus.ExpeditionManagement"/>
                    <Waifus v-else-if="navigationStatus == NavigationStatus.HeroMaker"/>
                    <InspectHero v-else-if="navigationStatus == NavigationStatus.InspectHero"/>
                </section>
            </section>
        </div>
    </section>
</template>
