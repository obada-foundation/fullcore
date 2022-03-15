<template>
  <div class="container">
    <div class="sp-component sp-component--half">
      <div class="sp-component-title">
        <h3>OBT NFT</h3>
      </div>

      <MintNftForm />
    </div>

    <div class="sp-component sp-component--half" v-show="loggedIn">
      <div class="sp-component-title">
        <h3>Transfer NFT</h3>
      </div>

      <TransferNFT />
    </div>

    <div class="sp-component" v-show="loggedIn">
      <div class="sp-component-title">
        <h3>NFT List</h3>
      </div>

      <NftList />
    </div>
  </div>
</template>

<style>
@media (min-width: 1366px) {
  .sp-component--half {
    flex-basis: 50%;
    width: 50%;
    padding-right: 2rem;
  }
}
</style>

<script>
import MintNftForm from "../components/MintNftForm"
import NftList from "../components/NftList"
import TransferNFT from "../components/TransferNFT"

export default {
  name: 'NFT',
  components: { MintNftForm, NftList, TransferNFT },
  computed: {
    loggedIn() {
      if (this._depsLoaded) {
        return this.$store.getters['common/wallet/loggedIn']
      } else {
        return false
      }
    },
  },
}
</script>
