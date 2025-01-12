import type { MetaFunction } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";

import { Button, InputGroup } from "~/components";

import styles from "./change-password.module.scss";

export const meta: MetaFunction = () => {
  return [
    { title: "Change password - Lopes Drive" },
    {
      name: "description",
      content:
        "If you forgot your password, fill up the form and proceed to change your password",
    },
  ];
};

export default function Index() {
  return (
    <div className={styles.container}>
      <main className={styles.card}>
        <h1>Change password</h1>

        <Form className={styles["change-password-form"]}>
          <InputGroup id="email" label="Email:" />

          <div className={styles["change-password-form__forgot-password"]}>
            <p>Already know your password?</p> <Link to="../login">Login</Link>
          </div>

          <Button type="submit">Change Password</Button>
        </Form>
      </main>

      <small>Copyright &copy; 2025 Matheus Lopes. All rights reserved.</small>
    </div>
  );
}
