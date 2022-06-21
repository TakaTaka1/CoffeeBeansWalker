<template>
    <header>
        <div id="hambergerMenu">
            <button type="button" class="menu-btn" v-on:click="open=!open">
            <i class="fa fa-bars" aria-hidden="true"></i>
            </button>
            <div class="menu" v-bind:class="{'is-active' : open }">
                <div class="menu__item">
                    <router-link to="/">TOP</router-link>
                </div>
                <div class="menu__item">
                    <router-link to="/BeansInfo">SETTINGS</router-link>
                </div>
                <div class="menu__item">
                    <router-link to="/BeansInfo">SHOPLIST</router-link>
                </div>                
            </div>            
            <!-- 参考 https://hakenblog.com/hamburger-menu/ -->
        </div>
    </header>
</template>

<script>
export default {
    el: '#hambergerMenu',
    data() {
        return {
            open: false
        }
    }
}

</script>

<style>
/*----------------------------
* メニュー開閉ボタン
*----------------------------*/
.menu-btn{
    position: fixed;
    box-shadow: 0  1px 0 rgba(136,136,136,1);
    top: 20px;
    right: 20px;
    z-index: 3;
    width: 40px;
    height: 40px;
    display: flex;
    border-radius: 3px;
    justify-content: center;
    align-items: center;
    background: #d3d3d3;
}

/*----------------------------
* メニュー本体
*----------------------------*/
.menu{
  position: fixed;
  top: 0;
  right: 0;
  z-index: 1;
  width: 20%;
  height: 30%;
  border-radius: 2px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: #555;
}
.menu__item{
    font-family: 'Caveat', cursive;
    width: 100%;
    height: auto;
    padding: .5em 1em;
    text-align: center;
    color: #fff;
    box-sizing: border-box;
}

/*----------------------------
* アニメーション部分
*----------------------------*/

/* アニメーション前のメニューの状態 */
.menu{
  transform: translateX(100vw);
  transition: all .3s linear;
}
/* アニメーション後のメニューの状態 */
.menu.is-active{
  transform: translateX(0);
}
</style>