import type { ActionFunctionArgs, MetaFunction } from "@remix-run/node";

import styles from "./landing.module.scss";
import { Button, InputGroup, Modal } from "~/components";
import { useState } from "react";
import { Form, useActionData } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "Lopes Drive" },
    { name: "description", content: "Welcome to the lopes drive app!" },
  ];
};

export default function Index() {
  const actionData = useActionData<typeof action>();
  const [isOpen, setOpen] = useState(false);

  console.log(actionData);

  return (
    <main className={styles.landing}>
      <div className={styles.landing__header}>
        <h2>My Drive</h2>

        <Button size="small" onClick={() => setOpen(true)}>
          Upload
        </Button>
      </div>

      <div className={styles.landing__content}>
        {resources.map(({ alt, src }) => (
          <ImageCard key={alt} alt={alt} src={src} />
        ))}
      </div>

      <Modal isOpen={isOpen} onClose={() => setOpen(false)}>
        <Form className={styles["landing__upload-form"]} method="POST">
          <h3>Upload your file</h3>

          <InputGroup label="File:" type="file" id="file" />

          <Button fullWidth type="submit">
            Upload
          </Button>
        </Form>
      </Modal>
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

export const action = (_: ActionFunctionArgs) => {
  return { success: true };
};
