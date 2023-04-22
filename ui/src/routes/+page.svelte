<script>
  import "@picocss/pico";
  import axios from "axios";

  let numbers;
  let count = 5;

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

<h1>로또번호 생성기</h1>
<p>로또번호를 생성합니다.</p>
<div class="grid">
  <input type="number" bind:value={count} />
  <button on:click={generate}>생성</button>
</div>
{#if numbers}
  <h2>생성된 로또번호</h2>
  <p>생성된 로또번호는 다음과 같습니다.</p>
  <article>
    {#each numbers as number}
      <div class="lotto">
        {#each number as num}
          <span>{num}</span>
        {/each}
      </div>
    {/each}
  </article>
{/if}

<style>
  .grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 1rem;
  }
  .lotto {
    display: grid;
    grid-template-columns: repeat(6, 1fr);
    grid-gap: 1rem;
    margin-bottom: 1rem;
  }
  .lotto span {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 3rem;
    height: 3rem;
    border: 1px solid #ccc;
    border-radius: 0.5rem;
    font-size: 1.5rem;
    font-weight: bold;
  }
</style>
