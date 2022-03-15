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
          <td>
            <router-link
                class="sp-blockdisplayline__height__link"
                :to="{ name: 'NFTDetails', params: { did: nft.id }}">{{ nft.data.usn }}</router-link>
          </td>
          <td class="sp-blockdisplayline__height">
            {{ nft.id.substr(0, 22) }}...
            <a href="#" class="sp-copy-icon" @click="copyStr(nft.id)">
              <span class="sp-icon sp-icon-Copy" />
            </a>
          </td>
          <td class="sp-blockdisplayline__hash">
            {{ nft.uri }}
          </td>
          <td class="sp-blockdisplayline__timestamp">
            {{ shortHash(nft.data.checksum) }}
            <a href="#" class="sp-copy-icon" @click="copyStr(nft.data.checksum)">
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
import { copyToClipboard, shortenHash } from '@/utils/helpers'

export default {
  name: 'NftList',
  created() {
    this.initNFTs()
  },
  methods: {
    copyStr(str) {
      copyToClipboard(str)
    },
    shortHash(string) {
      return shortenHash(string)
    },
    initNFTs() {
      this.$store
        .dispatch('obadafoundation.fullcore.obit/QueryGetNftsByAddress', {
          params: {
            address: this.currentAccount,
          },
          options: {
            subscribe: true,
            all: true,
          },
        })
        .then(function (resp) {
          console.log(resp)
        })
    },
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
      return this.$store.getters['obadafoundation.fullcore.obit/getGetNftsByAddress']({
        params: {
          address: this.currentAccount
        }
      })?.NFT ?? []
    },
  },
}
</script>
