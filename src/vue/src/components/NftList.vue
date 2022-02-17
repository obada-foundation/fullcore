<template>
  <div class="sp-box sp-shadow">
    <table class="sp-blockstable">
      <thead>
        <tr>
          <th class="sp-blockstable__height"><strong>USN</strong></th>
          <th class="sp-blockstable__height"><strong>DID</strong></th>
          <th class="sp-blockstable__hash"><strong>Metadata URL</strong></th>
          <th class="sp-blockstable__timestamp"><strong>Root Hash</strong></th>
        </tr>
      </thead>
      <tbody>
        <tr class="sp-blockdisplayline" v-if="!nfts || nfts.length === 0">
          <td colspan="3">No NFTs yet</td>
        </tr>
        <tr class="sp-blockdisplayline" v-for="nft in nfts" v-bind:key="nft.id">
          <td>{{ nft.data.usn }}</td>
          <td class="sp-blockdisplayline__height">
            {{ nft.id.substr(0, 22) }}...
            <a href="#" class="sp-accounts-list-item__copy" @click="copyStr(nft.id)">
              <span class="sp-icon sp-icon-Copy" />
            </a>
          </td>
          <td class="sp-blockdisplayline__hash">
            {{ nft.uri }}
          </td>
          <td class="sp-blockdisplayline__timestamp">
            {{ nft.data.checksum.substr(0, 12) }}...
            <a href="#" class="sp-accounts-list-item__copy" @click="copyStr(nft.data.checksum)">
              <span class="sp-icon sp-icon-Copy" />
            </a>
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
//import { copyToClipboard } from '@starport/vue/src/utils/helpers'

export default {
  name: 'NftList',
  created() {
    this.initNFTs()
  },
  methods: {
    copyStr(str) {
        const el = document.createElement('textarea')
        el.value = str
        document.body.appendChild(el)
        el.select()
        el.setSelectionRange(0, 999999)
        document.execCommand('copy')
        document.body.removeChild(el)
    // copyToClipboard(did)
    },
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
