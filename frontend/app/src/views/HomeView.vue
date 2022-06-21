<template>
    <div class="topContents">
        <div class='elems'>
            <div class='title'>
                <h1><label> Coffee Beans Walker</label></h1>
            </div>
            <div class="searchShopBtn">
                <input type="submit" value="&#xf002;" class="searchShopBtn btn--circle btn--circle-c btn--shadow fas">
            </div>
        </div>
    </div>

</template>

<style>

.elems {
    font-family: 'Caveat', cursive;
    font-size: large;
    text-align: center;
    width: 200px;
    height: 200px;
    position: absolute;
    top:0;
    right:0;
    bottom:0;
    left:0;
    margin:auto;
}

.searchShopBtn.btn--circle {
  width: 30px;
  height: 30px;
  padding: 0;
  background: #fff;
  border-radius: 50%;
  -webkit-box-shadow: 0 3px 5px rgba(0, 0, 0, .3);
  box-shadow: 0 3px 5px rgba(0, 0, 0, .3);
}

</style>