import type { MetaFunction } from "@remix-run/node";

import styles from "./landing.module.scss";
import { Button } from "~/components";

export const meta: MetaFunction = () => {
  return [
    { title: "Lopes Drive" },
    { name: "description", content: "Welcome to the lopes drive app!" },
  ];
};

export default function Index() {
  return (
    <main className={styles.landing}>
      <div className={styles.landing__header}>
        <h2>My Drive</h2>

        <Button variant="small">Upload</Button>
      </div>

      <div className={styles.landing__content}>
        {resources.map(({ alt, src }) => (
          <ImageCard key={alt} alt={alt} src={src} />
        ))}
      </div>
    </main>
  );
}

function ImageCard({ alt, src }: Readonly<(typeof resources)[number]>) {
  return (
    <div className={styles.file} key={alt}>
      <img className={styles.file__image} src={src} alt={alt} />
    </div>
  );
}

const resources = [...Array(5)].map((_, index) => ({
  src: `/assets/photo${index + 1}.jpg`,
  alt: `animal ${index + 1}`,
}));
