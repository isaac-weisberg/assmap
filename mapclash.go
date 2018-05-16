package main

func (localAssmap *assmap) squashLocalAssmapAgainst(assmap *assmap) {
	for _, asset := range assmap.Assets {
		localAsset := localAssmap.assetAt(asset.Path)
		if localAsset != nil {
			if localAsset.Modification != asset.Modification {
				asset.Version += localAsset.Version
			}
		}
	}
}
