<script>
  import "@picocss/pico";
  import axios from "axios";

  let numbers;
  let count = 1;

  const generate = async () => {
    const data = await generateLottoNumbers();
    console.log(data);
    numbers = data.numbers;
  };

  // /api/lottos API를 호출하여 로또번호를 생성하는 함수
  // fetch를 사용하여 API를 호출하고, 응답을 받아서 로또번호를 생성한다.
  const generateLottoNumbers = async () => {
    const data = await axios.get("/api/lottos?count=" + count);
    return data.data;
  };
</script>

<h1>로또번호 생성</h1>
<p>로또번호를 생성합니다.</p>
<button on:click={generate}>생성</button>
<input type="number" bind:value={count} />
{#if numbers}
  <h2>생성된 로또번호</h2>
  <p>생성된 로또번호는 다음과 같습니다.</p>
  {#each numbers as number}
    <div>{number}</div>
  {/each}
{/if}
