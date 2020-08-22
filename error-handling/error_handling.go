package erratum

import "log"

func Use(o ResourceOpener, input string) error {

	if resource, err := o(); err != nil {
		if _, is := err.(TransientError); is {
			if resource == nil {
				log.Println(resource)
			} else {
			}
			Use(o, input)
		} else {
			log.Println(err)
			log.Println(resource)
			return err
		}

	} else {
		resource.Frob(input)
		return resource.Close()
	}

	return nil
}
