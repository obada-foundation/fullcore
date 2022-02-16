<template>
  <div class="container">

    <div class="sp-type">
      <div class="sp-type__header sp-component-title">
        <h3>OBT NFT</h3>
      </div>

      <div class="sp-type__holder">
        <div class="sp-type-form__holder">
          <MintNftForm />
        </div>
      </div>
    </div>

    <div class="sp-type" v-show="loggedIn">
      <div class="sp-type__header sp-component-title">
        <h3>NFT List</h3>
      </div>

      <NftList />
    </div>

  </div>
</template>

<style>
.sp-type {
  margin-bottom: 3rem;
}

.sp-type .sp-type__holder .sp-type-form__holder {
  flex-basis: 50%;
  width: 50%;
}

@media (max-width: 600px) {
  .sp-type .sp-type__holder .sp-type-form__holder {
    flex-basis: 100%;
    width: 100%;
  }
}
</style>

<script>
import MintNftForm from '../components/MintNftForm'
import NftList from '../components/NftList'

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
  },
}
</script>
