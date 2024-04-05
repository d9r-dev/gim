<script lang="ts">
  import logo from "./assets/images/logo-universal.png";
  import { OpenFile } from "../wailsjs/go/main/App.js";
  import { SetText, GetText } from "../wailsjs/go/main/TextPieceTable.js";
  import {EventsOn} from "../wailsjs/runtime/runtime";

  let text = "";

  EventsOn('onFileRead', onOpenFile)

  function onOpenFile(response) {
    text = response.Data
  }

  function openFile() {
    OpenFile().then((response) => {
    })
  }

</script>

<main>
    <button class="btn" on:click={openFile}>Open File</button>
    <div>{text}</div>
</main>

<style>
</style>
