// example 1 : fetch data from variable

import type { NextPage } from 'next';
import Head from 'next/head';
import Style from '../styles/Home.module.css';

const products = [
  { id: 1, title: 'First Product' },
  { id: 2, title: 'Second Product' },
];

const Home: NextPage = () => {
  console.log('[HomePage] render: ', products);
  return (
    <div>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={Style.main}>
        <ul>
          {products.map((product) => (
            <li key={product.id}>{product.title}</li>
          ))}
        </ul>
      </main>
    </div>
  );
};

export default Home;