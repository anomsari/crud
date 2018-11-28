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
import '@polymer/iron-flex-layout/iron-flex-layout.js';
import '@polymer/paper-item/paper-item.js';
import '@polymer/iron-list/iron-list.js';
import './shared-styles.js';

class MyMatakuliah extends PolymerElement {
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
          margin: 24px;
          padding: 16px;
          color: white;
          border-radius: 5px;
          background-color: pink;
          box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 1px 5px 0 rgba(0, 0, 0, 0.12), 0 3px 1px -2px rgba(0, 0, 0, 0.2);

        }
        .kartu:hover{
          background-color: green;
          cursor : pointer;
        }
      </style>

      <div class="card">
        <div class="circle">1</div>
        <h1>Matakuliah</h1>

        <paper-input id="kd_mk" label="Kode Mata Kuliah" on-change="_change" readonly></paper-input>
        <paper-input id="mata_kuliah" label="Mata kuliah" on-change="_change"readonly></paper-input>
      <br>
        <!-- <label>NPM   : </label><label id="npmnya"></label><br>
        <label>Nama  : </label><label id="namanya"></label><br>
        <label>Kelas : </label><label id="kelasnya"></label><br> -->

        <br>
        <!-- masukin data dummy pake list -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <div class="kartu" npm$="[[item.npm]]" >
            <div> no    :[[_index(index)]]</div>
            <div> kd_mk   : [[item.kd_mk]]</div>
            <div> mata_kuliah  : [[item.mata_kuliah]]</div>
          </div>
          </div>
          </template>
        </iron-list>
        <!-- <p id="npmnya"></p>
        <p id="namanya"></p>
        <p id="kelasnya"></p> -->

      </div>
    `;
  }

  static get properties(){
    return{
      matakuliah:{
        type    :Object,
        // notify  :true,
        value:{
          kd_mk:"",
          mata_kuliah:"",
        }
      },
      //contoh bikin data
      mhs:{
        type: Object,
        value:[
          {
            kd_mk   :"1si2",
            mata_kuliah  :"Sistem Informasi",

          },{
            kd_mk   :"1ea2",
            mata_kuliah  :"Sistem Data",

         },{
            kd_mk   :"1pa2",
            mata_kuliah  :"Sistem Terdistribusi",

          },{
            kd_mk   :"1st2",
            mata_kuliah  :"Algortima Pemrograman",

          }
        ]
      }
    };


  }

  _change(e){
    let key = e.srcElement.id;
    this.matakuliah[e.srcElement.id]=e.target.value;
  console.log(this.matakuliah);
}

  _index(index){
    var a = parseInt(index+1);
     return a
  }
  _listChanged(e){

    if (e.detail.value){
      let matakuliah = e.detail.value;
      this.$.kd_mk.value = matakuliah.kd_mk;
      this.$.mata_kuliah.value = matakuliah.mata_kuliah;
    }
  }
}

  //
  // _npm(e){
  //   // let npm = this.$.npmnya;
  //   // npm.innerHTML = e.target.value;
  //   this.$.npmnya.innerHTML = e.target.value;// manggil by id
  // }
  // _nama(e){
  //   // let npm = this.$.namanya;
  //   // nama.innerHTML = e.target.value;
  //       this.$.namanya.innerHTML = e.target.value;
  // }
  // _kelas(e){
  //   // let npm = this.$.kelasnya;
  //   // kelas.innerHTML = e.target.value;
  //       this.$.kelasnya.innerHTML = e.target.value;
  // }
  //


window.customElements.define('my-matakuliah', MyMatakuliah);
