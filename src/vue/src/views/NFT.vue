<template>
  <div>
    <div class="container">
      <h3>OBT NFT</h3>
      <MintNftForm />

      <h3>NFT List</h3>
      <NftList v-show="loggedIn" />
    </div>
  </div>
</template>

<script>
import MintNftForm from "../components/MintNftForm"
import NftList from "../components/NftList"
  
export default {
  name: 'NFT',
  components: { MintNftForm, NftList },
  computed: {
    loggedIn() {
      if (this._depsLoaded) {
        return this.$store.getters['common/wallet/loggedIn']
      } else {
        return false
      }
    },
  }
}
</script>
