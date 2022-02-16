<template>
  <div>
    <table class="sp-blockstable">
      <thead>
        <tr>
          <th class="sp-blockstable__height"><strong>DID</strong></th>
          <th class="sp-blockstable__height"><strong>USN</strong></th>
          <th class="sp-blockstable__hash"><strong>Metadata URL</strong></th>
          <th class="sp-blockstable__timestamp"><strong>Root Hash</strong></th>
        </tr>
      </thead>
      <tbody>
        <tr class="sp-blockdisplayline" v-for="nft in nfts" v-bind:key="nft.id">
          <td class="sp-blockdisplayline__height">
            <router-link :to="/nfts/" class="sp-blockdisplayline__height__link"
              >{{ nft.id }}
            </router-link>
          </td>
          <td>{{ nft.data.usn }}</td>
          <td class="sp-blockdisplayline__hash">
            {{ nft.data.uri }}
          </td>
          <td class="sp-blockdisplayline__timestamp">
            {{ nft.data.root_hash }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style>
.option-radio > .button {
  height: 40px;
  width: 50%;
}
</style>

<script>
export default {
  name: 'NftList',
  created() {
    this.initNFTs()
  },
  methods: {
    initNFTs() {
      this.$store.dispatch("obadafoundation.fullcore.obit/QueryGetAllNftByOwner", {
        params: {
          owner: this.currentAccount,
        },
        options: {
          subscribe: true,
          all: true
        }
      }).then(function (resp) {
        console.log(resp)
      });
    }
  },
  computed: {
    currentAccount() {
      if (this._depsLoaded) {
        if (this.loggedIn) {
          return this.$store.getters['common/wallet/address']
        } else {
          return null
        }
      } else {
        return null
      }
    },
    loggedIn() {
      if (this._depsLoaded) {
        return this.$store.getters['common/wallet/loggedIn']
      } else {
        return false
      }
    },
    nfts() {
      return this.$store.getters['obadafoundation.fullcore.obit/getGetAllNftByOwner']({
        params: {
          owner: this.currentAccount
        }
      })?.NFT ?? []
    }
  }
}
</script>
