insert into public.company_info (short_name, full_name, phone, email)
values (
        'Фела',
        'Компания Фела',
        '+7(789)123-45-67',
        'info@example.ru'
    );
insert into public.page (id, parent_id, url, name, title, description)
values (
        1,
        null,
        '/',
        'Главная',
        'Услуги от компании Фела',
        'Много разных полезнух услуг от компании Фела'
    );
insert into public.content (page_id, tag, data)
values (
        '1',
        'main',
'<h1 class="text-center display-5 font-weight-normal m-4">
    Lorem ipsum dolor sit amet
</h1>

<div class="text-justify content">р
    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
</div>

<h2 class="text-center display-6 font-weight-normal">
    Lorem ipsum dolor sit amet?
</h2>

<ul class="ml-4">
    <li>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua;</li>
    <li>Nisl tincidunt eget nullam non;</li>
    <li>Quis hendrerit dolor magna eget est lorem ipsum dolor sit;</li>
    <li>Volutpat odio facilisis mauris sit amet massa.</li>
</ul>

<h2 class="text-center display-6 font-weight-normal mt-4">
    Nunc pulvinar sapien et ligula ullamcorper malesuada proin:
</h2>
'
    ),
    (
        '1',
        'bottomText',
'<div class="text-justify m-4 pl-4 pr-4">
    <p class="lead font-weight-normal text-center">
        Libero id faucibus nisl tincidunt eget. Leo a diam sollicitudin tempor id
    </p>
</div>

<div class="text-center w-100">
    <a class="btn btn-primary btn-lg" href="/contacts">Написать</a>
</div>'
    );