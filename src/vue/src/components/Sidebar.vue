<template>
  <SpSidebar v-on:sidebar-open="sidebarOpen = true" v-on:sidebar-close="sidebarOpen = false">
    <template v-slot:default>
      <SpLinkIcon link="/" text="Dashboard" icon="Dashboard" />
      <SpLinkIcon link="/trust-anchors" text="Trust Anchors" icon="Form" />
      <SpLinkIcon link="/nft" text="OBT NFT" icon="Form" />
      <SpLinkIcon link="/relayers" text="Relayers" icon="Transactions" />
      <div class="sp-dash"></div>
      <SpLinkIcon href="https://www.obada.io/standard/" target="_blank" text="OBADA Documentation" icon="Docs" />
    </template>
    <template v-slot:footer>
      <SpStatusAPI :showText="sidebarOpen" />
      <SpStatusRPC :showText="sidebarOpen" />
      <SpStatusWS :showText="sidebarOpen" />
      <div class="sp-text">Build: v0.3.8</div>
    </template>
  </SpSidebar>
</template>

<script>
export default {
  name: 'Sidebar',
  data() {
    return {
      sidebarOpen: true,
    }
  },
  computed: {
    hasWallet() {
      return this.$store.hasModule(['common', 'wallet'])
    },
  },
}
</script>
