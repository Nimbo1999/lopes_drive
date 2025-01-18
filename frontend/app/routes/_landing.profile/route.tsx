import type { MetaFunction } from "@remix-run/node";
import { Form } from "@remix-run/react";

import Styles from "./profile.module.scss";
import { Button, InputGroup } from "~/components";

export const meta: MetaFunction = () => {
  return [
    { title: "My Profile" },
    {
      name: "description",
      content: "This is the user profile page of the driver app",
    },
  ];
};

export default function Route() {
  return (
    <main className={Styles.profile}>
      <h2>My profile</h2>

      <Form className={Styles["profile__form"]}>
        <InputGroup id="username" label="Username:" />
        <InputGroup id="email" label="Email:" />
        <Button type="submit">Update profile</Button>
      </Form>
    </main>
  );
}
