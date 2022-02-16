<template>
  <div>
    <div class="container">
      <h3>OBT NFT</h3>
      <MintNftForm />

      <h3>Transfer NFT</h3>
      <TransferNFT v-show="loggedIn" />

      <h3>NFT List</h3>
      <NftList v-show="loggedIn" />
    </div>
  </div>
</template>

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
  }
}
</script>
