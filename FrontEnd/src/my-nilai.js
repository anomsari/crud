/**
 * @license
 * Copyright (c) 2016 The Polymer Project Authors. All rights reserved.
 * This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
 * The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
 * The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
 * Code distributed by Google as part of the polymer project is also
 * subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
 */

 import { PolymerElement, html } from '@polymer/polymer/polymer-element.js';
 import '@polymer/paper-input/paper-input.js';
 import '@polymer/paper-card/paper-card.js';
 import '@polymer/iron-flex-layout/iron-flex-layout.js';
 import '@polymer/paper-item/paper-item.js';
 import '@polymer/paper-item/paper-icon-item.js';
 import '@polymer/iron-list/iron-list.js';
 import '@polymer/iron-media-query/iron-media-query.js';
 import '@polymer/paper-fab/paper-fab.js';
 import './shared-styles.js';

class MyNilai extends PolymerElement {
  static get template() {
    return html`
      <style include="shared-styles">
        :host {
          display: block;

          padding: 10px;
        }
        .vertical{
          @apply -- layout-vertical;
        }
        .horizontal{
          @apply -- layout-horizontal;
        }
        .kartu{
          border-collapse: collapse;
          width: 100%;
          text-align: left;
          padding: 2px;
        }

        .kartu:hover{
          background-color: green;
          cursor : pointer;
        }
        .kartu-mobile{
          border-collapse: collapse;
          width: 100%;
          text-align: left;
          padding: 2px;
        }
        .kartu-mobile:hover{
          background-color: red;
          cursor : pointer;
        }

        .paper-fab[label=Y]{
        front-weight :400px;
        front-size   :200px;
        }

        .paper-fab[label=Z]{
        front-weight :400px;
        front-size   :200px;
        }
        .margin{
          margin-left 10px;
        }
      </style>


      <div class="card">
        <div class="circle">1</div>
        <h1>Nilai</h1>
        <paper-item class="horizontal">
        <paper-input id="npm" label="search" on-change="_change" readonly></paper-input>
        </paper-item>

        <br>
        <!-- <label>NPM   : </label><label id="npmnya"></label><br>
        <label>uas  : </label><label id="uasnya"></label><br>
        <label>uts : </label><label id="utsnya"></label><br> -->
        <iron-media-query query="(min-width : 641px)" query-matches="{{desktop}}"> </iron-media-query>

        <br>
        <template is="dom-if" if="{{desktop}}">
        <!-- masukin data dummy pake list/ tampilan desktop -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <paper-item class="kartu horizontal" kd_mk$="[[item.kd_mk]]">
          <!-- <div class="kartu" npm$="[[item.npm]]" > -->
            <!-- <div> no    :[[_index(index)]]</div> -->
            <div style="width :20%"> [[item.kd_mk]]</div>
            <div style="width :20%"> [[item.npm]]</div>
            <div style="width :10%"> [[item.uas]]</div>
            <div style="width :10%"> [[item.uts]]</div>
            <div style="width :20%"> [[item.total]]</div>
            <div style="width :20%"> [[item.ipk]]</div>
          </div>
          </paper-item>
          <!-- </div> -->
          </template>
        </iron-list>
        </template>

        <template is="dom-if" if="{{!desktop}}">
        <!-- masukin data dummy pake list/ tampilan hp -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <paper-item class="kartu-mobile horizontal" kd_mk$="[[item.kd_mk]]">
          <paper-fab label="[[_substring(item.kd_mk)]]">  </paper-fab>
          <!-- <div class="kartu-mobile" npm$="[[item.npm]]" > -->
            <!-- <div> no    :[[_index(index)]]</div> -->
            <div class="vertical margin">
            <div style="width :20%"> [[item.kd_mk]]</div>
            <div style="width :20%"> [[item.npm]]</div>
            <div style="width :10%"> [[item.uas]]</div>
            <div style="width :10%"> [[item.uts]]</div>
            <div style="width :20%"> [[item.total]]</div>
            <div style="width :20%"> [[item.ipk]]</div>

          </div>
          </div>
          </paper-item>
          </template>
        </iron-list>
        </template>

        <!-- <p id="npmnya"></p>
        <p id="uasnya"></p>
        <p id="utsnya"></p> -->

      </div>
    `;
  }

  static get properties(){
    return{
      // nilai:{
      //   type    :Object,
      //   // notify  :true,
      //   value:{
      //     kd_mk:"",
      //     npm:"",
      //     // uas:"",
      //     // uts:"",
      //     // total:"",
      //     // ipk:"",
      //         }
      // },
      //contoh bikin data
      mhs:{
        type: Object,
        value:[
          {
            kd_mk:"1si2",
            npm   :"51418634",
            uas  :56,
            uts :86,

          },{
            kd_mk:"1ea2",
            npm   :"51418997",
            uas  :54,
            uts :75,

          },{
            kd_mk:"1pa2",
            npm   :"51418462",
            uas  :42,
            uts :54,

          },{
            kd_mk:"1st2",
            npm   :"51418451",
            uas  :85,
            uts :54,

          }
        ]
      }
    };


  }
  _substring(s){

    var str = s.substring(1, 3);
    var sub = str.toUpperCase();
    return sub
  }
  _change(e){
    let key = e.srcElement.id;
    this.nilai[e.srcElement.id]=e.target.value;
  console.log(this.nilai);
}

  _index(index){
    var a = parseInt(index+1);
     return a;
  }

  _listChanged(a){
    if (a.detail.value){// cek datanya atau engga

    var nilai = a.detail.value;
    let uts = nilai.uts;
    let uas = nilai.uas;
    let total = (nilai.uts+nilai.uas)/2;
    let ipk ="";

     if ( total> 80){
      ipk = "A";
    }else if ( total> 70){
      ipk = "B";
    }else if ( total> 60){
      ipk = "C";
    }else if(total > 50) {
      ipk = "D";
    }else{
      ipk = "P";
    }
    // nilai.value = nilai.ipk;



      this.$.kd_mk.value = nilai.kd_mk;
      this.$.npm.value = nilai.npm;
      this.$.uas.value = nilai.uas;
      this.$.uts.value = nilai.uts;
      this.$.total.value = total;
      this.$.ipk.value = ipk;

    }
  }
}

  //
  // _npm(e){
  //   // let npm = this.$.npmnya;
  //   // npm.innerHTML = e.target.value;
  //   this.$.npmnya.innerHTML = e.target.value;// manggil by id
  // }
  // _uas(e){
  //   // let npm = this.$.uasnya;
  //   // uas.innerHTML = e.target.value;
  //       this.$.uasnya.innerHTML = e.target.value;
  // }
  // _uts(e){
  //   // let npm = this.$.utsnya;
  //   // uts.innerHTML = e.target.value;
  //       this.$.utsnya.innerHTML = e.target.value;
  // }
  //


window.customElements.define('my-nilai', MyNilai);
